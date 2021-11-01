# chapter 1
data-intensive vs computer-intensive -- CPU power isn't the limiting factor, instead it's the amount, complexity, and rate of change of the underlying data. 

goal is to achieve a system that is reliable, scalable, and maintainable. 

what is a data system? any system or service that stores data and provides an interface for accessing that data. Could be a database, a message queue, a cache. 

fault vs failure: fault is when a component deviates from it's specification, whereas failure is when the system as a whole stops providing the service to a the user. 

## Reliability

**reliability**: system should continue to work correctly even in the face of faults like hardware errors, software errors, or human errors. fault-tolerance, resilience. 

critical bugs are often due to poor error handling, so reliability can be increased by intentionally causing faults in the system (instead of waiting for faults to happen). 

**hardware faults**: disk crashes, faulty ram, power outage, human unplugs something. 10k disks, 1 disk dies per day. 

MTTF: mean time to failure. Hard Disks = 10-50 years. 

Moving from hardware redundancy to software fault-tolerance techniques in preference or addition to hardware redundancy. 

## Scalability

**scalability**: as the system grows in data volume, traffic, or complexity it should be able to deal with that growth -- the ability for a system to cope with increased load. 

"if the system grows in a particular way, what are the options for coping with the growth"

important to be able to describe the load of the system in terms of load parameters: examples include requests per second, ratio of reads to writes, simultaneous active users, cache hit rate, etc. 

performance - latency vs response time. Response time is from the client perspective, latency is duration request is waiting to be handled. 

1. Be able to describe the load via a load parameter. 
2. Be able to measure performance in relation to the load parameter. 
3. Determine how to cope with increased load (scale up, scale out, re-architect)

**tail latency**: describes the highest end of your performance (eg 99th percentile of all users experience N response time)

## Maintainability

**maintainability**: people should be able to work on the system productively over time (fixing bugs, investigating failures, adapting it)

**accidental complexity**: complexity that is not inherent in the problem that the software solves but just arises from the implementation 

create good abstractions to pay down accidental complexity. 

apply Agile and TDD and refactoring to make systems easy to change in future. 

# Intro Lecture Notes 

distributed systems is basically "what happens when you need to scale beyond a single node" -- examining the tradeoffs, eg buy throughput but pay in consistency 

Amazon led the charge into microservices -- Meera: Looks like this is the Werner Vogels talk that Oz mentioned https://queue.acm.org/detail.cfm?id=1142065 

Fundamental distributed systems tradeoff: Replication and partioning in exchange for consistency. 

https://danluu.com/postmortem-lessons/

We don't have good systems for testing configuration changes, and thus a lot of outages are caused by configuration changes breaking something in an unexpected way. 

NTP is a hack, it takes time to communicate the time. 

be eternally pessimistic about your ability to write bug-free software, no amount of testing is never enough. 

is automatic failover a good thing? depends on the context. Most of the time, probably good, but if the system makes a mistake and thinks you are down but you aren't, what's the cost of switching to the replica? 

TODO: read the recent FB postmortem, read the Github 2012 post-mortem, read the complex systems memo, also watch this Bryan Cantrill talk https://www.youtube.com/watch?v=30jNsCVLpAE

### Rich Hickey Keynote on systems 

https://www.youtube.com/watch?v=ROor6_NGIWU


# Distributed Data

scaling out can be useful for: scalability, fault tolerance/high availability, and latency reduction. 

scaling up == shared-memory architecture. Also a shared-disk architecture (but you run into contention and locking limit)

shared-nothing architecture aka horizontal scaling or scaling out. 

replication: copying data across nodes to provide redundancy in case of failure and improve performance 
partitioning: splitting dataset into subsets across nodes. 

## Replication 

why replicate data? 
- to keep data close to users geographically and reduce latency 
- increase availability in case of failures 
- scale out to increase performance 

You can replicate using different strategies: **single-leader, multi-leader, or leaderless** -- and **synchronous or asyncronous**. 

**leader based replication**: you have a single node that is the leader and all writes must go through the leader. other nodes are followers or read replicas - leader sends writes to the followers. This is built into postgres. 
**Sync vs async replication** downsides to both -- synchronous is nice because it increases durability of write but could slow down performance (if follower is unavailable). async is nice because increases availability but decreases durability of writes / could lead to losing data. 

Adding followers happens basically by copying a snapshot of the leader and then the follower requesting all data starting from where the snapshot ends. 

Leader failover (when leader is unavailable and a new leader is chose) can be tricky: hard to determine for sure if leader has failed, also have to get cluster to agree on new leader. 

Replication log strategies: WAL is good because it's the data being written but it's storage implementation specific. Logical log similar idea but less coupled to storage allowing for backward compatibility (eg leader and follower could run different versions of software)

**eventual consistency** an effect where data read from a follower may be outdated because of the replication lag between it and the leader. 

**read after write consistency** guarantee that if user reloads page, they will see the updates they have submitted even in an eventually consistent system. 

**monotonic reads** guarantee that reads won't go backward in time (eg if you read something that happens, you will not read from a follower that has an older copy of the DB that what you've already read)

**consistent prefix reads**: guarantee that if sequence of writes happens in a certain order then a reader will see them appear in the same order. 

Why would you want a multi-leader setup? 

- could help if you have multiple data centers, you could have clusters that have single leader, and across the datacenters have a multi leader setup. 

If you have more than one leader, this can lead to conflicts: 

- best way to resolve a conflict is to avoid it in the first place (by having a single leader, or routing requests from a single user through the same leader)
- make unit of change really small to reduce likelihood of conflicts. 
- last write wins or oldest replica wins -- both are unambiguous but will lead to data loss
- could record the conflict and then ask the user to resolve it. 

