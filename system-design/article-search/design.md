# Problem statement
Add search to the [QA site](../q-and-a/design.md). 

# Requirements
- ability to search through questions and answers returning results ranked by relevancy, votes, views, and recency. 
- spellcorrect yes, synonyms no.
- (bonus) analytics on search queries, clickthrough rates on the ranked results. 
- (bonus) think about rate limiting.

# SLOs
- 95% searchable within 10 seconds after posting, 100% after 1 day. 
- 99% of requests should return within a second. 

# Out of scope
- typeahead / autocomplete out of scope 
- cost comparison to search as a service.
- focus of the problem is not inventing search features. 
- do not want hot recent searches on the homepage.

# Load Estimation 
expected search volume: 200M a month, 100% YoY growth
1st year 200M => 80rps if spread evenly but can spike to 500-1k rps 
2nd year 400M => 160rps 
3rd year 800M  => ~300rps

See the original [QA site](../q-and-a/design.md) for data model of the question and answer data set. 

# Ways to support search

There are many ways we could possibly support search, explored below.

## Postgres 

Postgres supports full-text search as of version 9.5 ([docs](https://www.postgresql.org/docs/14/textsearch-intro.html)) via indexing documents by their tokens and lexemes. 

In our usecase, we could start by adding a column to our question and answers tables to automatically generate searchable indexes for the specific columns we need to search eg

```
ALTER TABLE questions
    ADD COLUMN search_index tsvector
               GENERATED ALWAYS AS (to_tsvector('english', coalesce(title, '') || ' ' || coalesce(body, ''))) STORED;
```
```
ALTER TABLE answers
    ADD COLUMN search_index tsvector
               GENERATED ALWAYS AS (to_tsvector('english', coalesce(body, '') )) STORED;
```

and then add a GIN index to these columns ([docs](https://www.postgresql.org/docs/14/textsearch-indexes.html)) -- GIN stands for Generalized Inverted Index, basically a token:[locations] map.

```
CREATE INDEX question_search_idx ON questions USING GIN (search_index);
CREATE INDEX answer_search_idx ON answers USING GIN (search_index);
```

Postgres supports ranking based on normalized weights, and custom ranking functions so we could return results that were ranked by a tuned combination of question and answer relevancy, vote count, and recency. It also supports returning "headlines" (matching parts of documents) which could be returned as part of the search result (although this can be slow because it doesn't use the index to generate the headline, it searches the whole document).

### Would it work?

**Storage:** Let's assume the tsvector column doubles the size of each row in the average case, because it removes stop words and common words but adds weights. Before the tsvector, we were expecting ~125GB per year of storage required. Taking into account the size of the GIN indexes as well, let's round up to expect around 300 GB per year. This is still do-able for multiple years on a single disk.

**Throughput:** Let's assume that a single search query will require 5-10 IO operations (following our GIN index across multiple tables). We know that we expect load to potentially spike to 1k at times during the first year. Thus, at peak load we might require disk throughput of 10k IOPS, which exceeds what we might be able to achive with a general purpose SSD, but would be achievable with an upgraded SSD (eg with provisioned IOPS) and/or introducing read-replica nodes to handle part of the search load. [Ref: AWS EBS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS). 

NB another approach here would be combining the data to be searched into a single TS Vector in a separate table. This would probably speed up the query significantly by reducing the number of IOs required.

### Tradeoffs 

Pros
* simplicity: no additional components required (eg team doesn't have to learn how to use a new tool like Elasticsearch)
* cost: only additional cost would be adding read replicas for postgres, no need to pay for other additional components like a managed ES cluster or 3rd party search service. 
* you get search indexing for free basically because postgres is already the primary data store 

Cons 
* although highly configurable, you have significantly less configurability than if you used a tool like Elasticsearch

## Elasticsearch / Solr

There are open source software projects built on Lucene that are specifically focused on search such as Elasticsearch (or a successor AWS's OpenSearch), and Apache Solr. 

### Tradeoffs

Pros
* extremely configurable, much more flexible than postgres text search.
* extremely performant 

Cons
* can be complex, new tool to learn if you don't have in-house expert already 
* added cost for cluster of servers. 
* would require sending data to the search cluster for indexing

## 3rd Party Search Service

Other companies provide [search as a service](https://en.wikipedia.org/wiki/Search_as_a_service) SaaS products, such as [Algolia](https://www.algolia.com/). 

### Would it work?

Build vs buy tradeoff. This is a viable option depending on team structure and budget. It could be a cheaper way in the short term to test out search without committing too many internal resources, with a long-term plan to build search in-house down the road.

### Tradeoffs 

Pros
* offloading complexity of search from your team / simplicity => search is an API integration away. 

Cons 
* giving up control over your availability and latency to a third party. If they have an outage, your customers suffer and your only recourse is legally enforcing the SLA that you have with them. Algolia provides 5-nines but you pay for it. 
* vendor lock -- if their prices go up, you have to eat the cost or eat the cost of migrating to a new provider or revisit the build vs buy decision. 
* cost: $1 per 1k search requests. 

# Cost Comparison

NB that this is the estimated cost we'd be adding to our existing cost in the original design. 

|Tool|Notes|Estimated Additional Cost|Source|
|---|---|---|---|
|Postgres|Let's assume we use a provisioned SSD for 5k IOPS and a read replica. So two nodes backed by SSDs with provisioned IOPS|($0.125 per GB per month * 500 GB * 2 nodes) + ($0.065 * 2000 addtional IOPS * 2(we get 3k for free)) = $125 storage +  $260 IOPS = $385/month |[AWS EBS Pricing](https://aws.amazon.com/ebs/pricing/)|
|ElasticSearch/Solr|Let's assume we use AWS OpenSearch which is ES in disguise, assume 2 large search nodes, with 500GB EBS volumes|$0.335/hr * 2 nodes * 24 hours * 30 days + 500GB * $0.135 * 2 = 482 compute + 135 storage =  $617/month|[AWS OpenSearch Pricing](https://aws.amazon.com/opensearch-service/pricing/)|
|Algolia/3rd Party|Algolia charges based on record count and search count. We expect to have to index all questions and answers, so annually we will have 25M records per year and 200M search requests|Unless I'm misunderstanding their pricing it seems like it's $1 per 1k records and 1k search requests. This would mean approx $2k per month in record costs and approx $15k per month in search costs. Complete non-starter, even with volume discounts. I must be misunderstanding something about their pricing.|[Algolia Pricing](https://www.algolia.com/pricing/)|

# Final Recommendation 

Since the data is already in postgres and we have the flexibility to vertically and/or horizontally scale our database cluster to support search, it seems like the least complex and most cost effective option would be to use Postgres as a starting point for support full-text search. 

If a new requirement arises that cannot be met by Postgres's full-text search capabilities or qualitatively we find that the search result relevancy is lacking, we could consider migrating to an open-source search engine such as Elasticsearch or Solr where there are many more knobs to turn with how documents are scored and how text is analyzed / tokenized.