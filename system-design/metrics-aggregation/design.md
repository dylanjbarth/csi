# Problem Statement 

We'd like to create a service that will allow us to collect performance data into a unified dashboard from our different services. 

# Functional Requirements
- should come with default metrics, but users should also be able to define their own. 
- should be able to review metrics up to 1 minute in granularity (not sub-minute). 
- metrics should be filterable / queryable in a simple way in a dashboard, but sophisticated search not required

# Non-functional Requirements
- must support all systems we've designed previously (and scale to support new services and growth in future).
- must be comprehensive and provide observability (unclear exactly how many metrics per service)
- availability should be high but not highest since this is not customer facing. 
- it's acceptable to drop some data. 
- it's acceptable to have more granular recent data vs historical data. 
- must support many geographic regions

# Out of scope

- Alarms / pager duty -- but should be extendable to this
- logs and traces for the metric collector itself. 

# SLOs

- users should be able to see most recent data within 1 minute. 
- availability target of 3 nines: 99.9% allowing for a minute or so of downtime per day ~8 hours per year. 

# Load Estimation 

## How big is a metric

In order to estimate load, we first need to understand what schema our core system object will have. We know that it requires some sort of service identifier, the name of the metric, a numeric value for the metric, a unit enum, and probably some additional metadata like tags. 

```
{
  namespace: varchar 50 bytes   # this identifies the broader service the metric is a part of 
  metric: varchar 100 bytes  
  value: int64 8 bytes  
  unit: enum 8 bytes
  timestamp: 8 bytes.
  tags: [  # up to 10, this allows customization for filtering and querying
    {
      name: varchar 50 bytes
      value: varchar 50 bytes
    }
  ]
}
```
Given this schema, let's assume a metric can reach a maximum size of 1500 bytes (rounding up to create buffer room for future fields). 

## Number of Services

Currently, ShadyCorp has the following services: 

* a Q&A site with search 
* an Image Hosting service 
* a URL Shortening service 

For each service, let's assume that we are going to want to collect a variety of performance related metrics for each sub-component of the service (eg CPU utilization, memory utilization, disk IO, network IO, by default plus many service specific custom metrics created by individual engineering teams). 

Each service has a variable number of components, but in general we can estimate that there is going to be at least 5 and in most cases no more than 10 total components (although individual components such as the database server or api servers may have many nodes -- thus they will send the types of metrics, but the total number of metrics would increase as they are scaled up).

To make this simple, let's assume that we limit the total number of custom metrics per namespace/service to 1000. This gives engineering teams a lot of flexibility to send metrics they care about for their specific use cases (this works out to 100 metrics per component, which is a ton!)

## Current Load

To account for growth, let's assume that within the first 2 years we want to be able to support 10 services, 1000 metrics per service, and that each metric can be updated once per minute. 

### Writes

||Total Metrics Per Minute | Metrics per second | Metrics Per Year |
|---|----|---|---|
|Volume|10 services * 1000 metrics = 10e3 metrics/minute| 100e2/60 => 170 metrics/second | 170 * 86.4e3 * 365 = ~5B |
|Size/Storage (1.5KB / metric) | 15MB/min |  250 KB/s | 7.5 PB / year *see note | 
|Bandwidth||250KB/s|||

Note: we know that we don't need to keep high resolution data indefinitely. We could compress older metrics to reduce our storage significantly. Eg instead of keeping minute level granularity, we could store the average, median etc for each hour, or even for the day. This would enable us to look at broad historical trends, and reduce our storage costs significantly.

If we rolled up metrics by the hour, 7.5 PB becomes 7.5 PB / 60 => 125 TB / year.

If we rolled up metrics by the day, 7.5 PB becomes 7.5 / 60 / 24 => 5.2 TB / year.

If we assume that we only need high resolution data for the past 2 weeks, we know that our total set of queryable data at any given time will be: 170 * 86.4e3 * 14 ~ 200M records, 300GB. 

### Reads

We have a small team, ~100 engineers, and we expect reads to be distributed across our namespaces/services relatively evenly. If each team member accesses the dashboard a few times a week, we would expect something like 20 active daily users. 

Let's assume that each session lasts 10 minutes they are looking at a dashboard and accessing 1000 metrics over this period.

Our API can return the data required to load the dashboard as a simple array of int64s to reduce the bandwidth required. An example request might look like: 

`fetchMetrics(namespace: str, metricUCK: str, start, end, filters)` where start and end are timestamps, and filters are pairs of tag key value pairs that must match (eg to filter to a specific component). A response would look like this: 

```
{
  "namespace": ...
  "metric": ...,
  "values": [value, value, value, etc]  # if we limit time ranges to the past two weeks, and assume we return values in 5 minute intervals, the max len of this array is (12 5 min intervals in an hour, 288 intervals in a day, 4032 in a week, * 8 bytes ~ 32 KB)
}
```

||Total Reads Per Day | Reads per second per session |
|---|----|---|
|Volume|1000 per session * 20 = 20e3 | We assume a user session lasts 10 minutes and they access 1k metrics. This gives us a more realistic estimate of read load than dividing total reads per day by number of seconds in a day. 1000/10 = 100 metrics/min / 60 = ~1.5 metrics per second.  | |
|Bandwidth||~32KB / read => 50KB/s|||

# Storage

Based on our load estimation, we need to be able to support a relatively high write throughput and a low read throughput. We also know that our working set of data (for the past 2 weeks), is small enough to fit on a single node (~300 GB, 200M records, 170 writes per second). We can add an asynchronous replica to increase availability in case of failover, although given our availability SLO, this probably isn't necessary. As long as we have regular backups, we should be able to restore from snapshot in enough time to meet our annual three nines SLO. 

TODO ERD

metric: 

id



tags


# Components 

A write API that writes directly to the DB. (could add a queue here? this would allow us to recover from DB failure and throttle during peak load)
DB w/o any failover. 
Aggregator service that reduces older data points and re-writes to metric table. 
Read api to handle read requests.

# Thoughts on this exercise

- I really struggled with this SD exercise because we didn't get clear ideas of scale from our PM, so I really spent a lot of time waffling about, trying to estimate the load on the system. If one assumption changes upstream, it cascades down to affect every other little assumption which can get exhausting. 
- could there be a better storage solution here? the metrics are immutable, so maybe just writing to disk? 