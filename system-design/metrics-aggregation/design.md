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

To consider: 

- how big is a metric?
- is there a limit to the number of metrics we allow per service? 
- if we get one per minute, how many will we get over time?
- what is our bottleneck going to be for writes? bandwidth? storage? 
- for our reading use case, what storage should we use? 
- what will our read bottleneck be? bandwidth? diskio? compute? 

## A single metric

In order to estimate load, we first need to understand what schema our core system object will have. We know that it requires some sort of service identifier, the name of the metric, a numeric value for the metric, a unit enum, and probably some additional metadata like tags. 

```
{
  namespace: varchar 50 bytes   # this identifies the broader service the metric is a part of 
  metric: varchar 100 bytes  
  value: int64 8 bytes  
  unit: enum 8 bytes
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

If we assume each component is going to send 5 metrics by default, and we assume that the engineering teams create an average of 15 custom metrics, then we can expect each component to send 20 metrics. 

## Current Load

Ignoring growth for now, let's attempt to calculate the load we might expect on the system given our current estimates. 

3 services
10 components per service
20 metrics per service 

### Writes

||Total Metrics Per Minute | Metrics per second | Metrics Per month | Metrics Per Year |
|---|----|---|---|---|
|Volume|3 services * 10 components * 20 metrics = 600 metrics/minute|10 metrics/second| 86.4e4 / day * 30 = ~25M | ~320M|
|Size/Storage| x1500 bytes = 900KB|15KB|~40GB|~500GB|
|Bandwidth||15KB/s|||

### Reads

We have a small team, ~100 engineers, and we expect reads to be distributed across our services relatively evenly. If each team member accesses the dashboard a few times a week, we would expect something like 20 active weekly users. Let's assume that in each session they are looking at either individual metrics at various levels of granularity (past hour, past day, past week), or they are looking at a dashboard showing multiple metrics at the same level of granularity. Thus during an active session, we might need to fetch quite a bit of data to populate these dashboards. 

Let's assume in a single session, a user looks at all 20 metrics for each component (max 10) of a single service, so 200 metrics total. Our API can return the data required to load the dashboard as a simple array of int64s to reduce the bandwidth required. An example request might look like: 

`fetchMetrics(namespace: str, metricNames: str[], start, end, filters)` where start and end are timestamps, and filters are pairs of tag key value pairs that must match. This would allow the client to fetch metrics for multiple dashboards at once. A response would look like this: 

We don't need to load minute level granularity for wider time ranges. For example, if the user is looking at the past 30 minutes, we could show them average data points at 5 minute intervals instead, to reduce the amount of data we are actually sending to the client. This would mean 6 data points per metric per 30 minute increment. We could also further improve load times for the full dashboard when looking at broader time ranges by reducing granularity further (10-15 minute averages, etc) and reducing the amount of data returned in a request -- we could send the metrics back as arrays of data points instead of the full metric. 

```
{
  "<namespace>": {
    "<metric>": [value, value, value, etc]  # if we limit time ranges to the past two weeks, and assume we return values in 5 minute intervals, the max len of this array is (12 5 min intervals in an hour, 288 intervals in a day, 4032 in a week, * 8 bytes ~ 32 KB)
  }
}
```

For the purposes of this exercise, with this schema for the response body, let's assume that we can set some sensible constraints on the total possible response size by bounding: 
- the total number of metrics in a single request
- the minimum interval between data points (as the different between start and end gets larger, so does the interval between individual data points). 

## Future Load

If we account for growth of existing services and introduction of future services, let's assume that our load will grow 100% year over year for the next three years. 

Using the future value formula, this means we can expect the following upper bounds: 
||1st Year|2nd Year|3rd Year|
|---|---|---|---|
|Storage|500GB|1TB|2TB|
|Write Bandwidth|||
|Read Bandwidth|||

Although this doesn't take into account the fact that we can expire some data... or roll it up in aggregations. 

# Thoughts on this exercise

- I really struggled with this SD exercise because we didn't get clear ideas of scale from our PM, so I really spent a lot of time waffling about, trying to estimate the load on the system. If one assumption changes upstream, it cascades down to affect every other little assumption which can get exhausting. 


- # services are the ones we've designed so far, double in the next year. 
- how many metrics per service? 
  - want to be comprehensive, need observability. 
- what level of granularity do we want on the metrics? 
  - 1 minute intervals. smallest interval 
- auto collect some metrics, specific service can add their metrics too. 
- "what fraction of requests are we getting, which are failing" 
- will appear in some dashboard where they can click around, filter by service, machine, not sophisticated search. 
- retention policy on data: 
  - not strict, but acceptable to have more recent data than historical data
  - would be great to aggregate - rollups.
- accuracy: acceptable to drop some data. 
- SLA for when data is up to date: should be low, could be a minute before they see most recent metrics. 
- Availability high, but not customer facing -- useful for our debugging purposes. 
- out of scope: alarms (but should be extendable to this). 
- budget - high priority: small team of engineers for 1 year. 
- lots of geographic regions. 
- how many users? all the engineers interacting with this frequently, at least a few times a week. ~100 people on the team. 