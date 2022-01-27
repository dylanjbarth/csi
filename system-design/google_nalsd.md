Principles: 
-  Start small (eg 1 machine) and identify problems, iterate from there 

# google adwords example

start by determining your SLO (time to load results and age of results): 

1. 99.9% of queries complete in < 1 second
2. 99.9% of queries read data that is <5 minutes old

understand the scale of the system: 

1. Number of search requests per second: 500k 
2. Number of clicks per second: 10k (2% of ads served)

Create an understanding of data size: 

- 1 log is ~2kb, 500k logs per second => 500,000 * 2,000 bytes => 1,00,000,000 bytes => 1M kb => 1,000 MB => 1 GB / second * 86,400 GB per day => 86.4 TB per day. 
- 2% of this in click logs per day => 8.64 / 5 = ~2TB extra per day in click logs. 

this helps justify changes to the design: 
- we know we need more than 1 machine 
- we know that storing this on disk is going to too slow (we'd need a ton of disks) 

think about how data needs to be transformed to fulfill the requirements

in this case, we have query_logs, and click_logs that need to be joined together. 

query log: { 
    timestamp
    search_term
    ad_ids: []
    query_id 
}

click_log { 
    timestamp
    ad_id
    query_id 
}

our goal is for users to be able to see what the CTR is per search term, so we need to join click_logs with query_logs to get: 

1. Total number of ad impressions per search term. 
2. Total number of clicks per search term. 

we have way more query logs than click logs, so if we can transform query logs to be rapidly searchable by query id, we can match click logs to query logs. 

if data has to pass between machines, calculate the network bandwidth needed. 
understand storage needs to store the transformed data. 


issues, what if an ad is quite popular, and we end up storing a lot of data in the KV store in the set of queries and timestamp clicked?

some random thoughts / questions: 

- when thinking about how to put a system like this together is it better to ignore the 'thing" that it is (eg the lego brick, like a load balancer, a specific database technology) and instead just focus on basic data structures, eg imagine we could store a large key value store here. 