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

**eventual consistency** an effect where data read from a follower may be outdated because of the replication lag between it and the leader. if all writes stopped, you will get the current answer after some indeterminate amount of time. 

**bounded staleness**: making eventual consistency stricter by enforcing a maximum of replication lag. 

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

# Partitioning 

aka sharding - breaking the entire dataset into chunks that are spread across nodes. 

why: scalability, allows you to distribute query load across many processors. 
goal is to spread load evenly, not doing so leads to skew and hotspots. 

rule of thumb: when to shard - only when you hit physical limitations. 

## partition strategies

**key range partitioning**: think encyclopedia volumes. splitting your primary keys by range evely. keys spaced based on distribution of your data so may not be even. keys are sorted within partition which makes range scanning easy. be careful of hot spots: eg if key is a timestamp, that creates a heavy workload for the day on a single partition. Be sure to pick a primary key that will distribute the load evenly. 

**hash partitioning**: must be fast, deterministic, and independent but doesn't need to be cryptographically strong. Distributes keys evenly. each partition gets a range of hashes. Impossible to do efficient range queries with this strategy. 

Cassandra use compound key: first one is hashed, second one is how it's sorted on disk, so you get even distribution and range querying. 

Even with key hashing, workload skew can happen due to a super popular key (eg celebrity ID) - falls to application to identify and reduce this skew. Eg adding a random number to beginning or end of key to split writes evenly. But this leads reads to having to do more work (query all keys and combine so more book keeping). 

## handling secondary indexes

**secondary indexes**: data that doesn't identify the record uniquely but is used as a way to search for occurrences, eg articles containing foo. 

**document based partitioning**: aka local index. store the secondary index within each partition as a lookup by the primary key. This means you have to search across partitions. known as "scatter/gather" because you have to query all partitions and then join together. you can try and choose partitioning scheme so that secondary indexes are clustered together but not always possible. 

**term-based partitioning**: aka global index. you use key range partitioning on the index itself. this can make reads more efficient because you hit a single partition that has the term you want, but writes are slower because a single document write could update many partition indexes. Updates to these indexes are therefore normally async. 

## rebalancing partitions

**rebalancing**: moving load from one node in the cluster to another. happens when you need to more CPUs, dataset increases, or nodes get replaced. 

**don't use hash(key) mod N**: this is because if the number of nodes changes, most keys will have to change nodes. we want to move data around as little as possible when we have to rebalance. 

**fixed number of partitions**: if you create many more partitions than nodes, you can logically shift a partition to a new node when you add one and not have to worry about re-partitioning. new nodes can steal a partition. While rebalancing, old partition is used until it's complete. works best with key hashing partitioning strategy. 

**dynamic partitioning** works better with key-range partitions. when range exceeds configured size, it's split in half. if it shrinks below, it gets merged. advantage is that partitions adapt to total data volume. This can also work for hash based partitioning. 

**proportional to nodes**: partition size grows proportionally to dataset size, but when you add node partitions shrink again

## automatic or manual rebalancing 

automated is convenient because of less work but can be unpredictable, so many systems tee it up but have human in loop. 

## request routing 

1. **node forwarding** clients contact any node and node services or forwards
2. **routing tier**
3. **client aware**

most systems rely on coordination service to keep routing components informed. Nodes register with service discovery system, eg ZooKeeper and it is the authoritative mapping to nodes, zookeeper then updates routing layers. 

another approach is gossip protocol, where nodes talk to each other. 

## Misc notes 

- consistent hashing: way to distribute load evenly across internet wide system of caches such as CDN. 

"fundamentally we want to partition based on load not based on key"

# Consistency & Consensus 

**consensus**: getting all nodes in a system to agree on something in spite of network faults or process failures. 

guarantee tradeoff: eventual consistency is hard to work with as an app developer because it's a leakier abstraction, but there is a tradeoff in performance. In general, the stronger the guarantee is the worse the performance or less fault tolerant a system is. 

**linearizability**: aka atomic consistency, strong consistency, immediate consisteny, or external consistency -- an abstraction or guarantee that makes it appears as if it were the only copy of the data and all operations are atomic. Thus application doesn't have to think about any of the complexity of the distributed system. It's a recency guarantee. 

**serializability**: guarantee that transaction behave as if they had executed in some serial order (isolated)

## Use cases for consensus: 

* leader election and locking: only want one leader and to elect need a lock. 
* constraints / uniqueness guarantees: need all nodes to agree on the value -- this is the atomic commit problem. 
* avoid race conditions across nodes (example of an async image resize)

**CAP theorem**: either consistent or available when partitioned -- if you requir linearizability and some replicas disconnect, you are down until they come back up. if you don't require it, you can be more fault tolerant but behavior is not linearizable. 

linearizability is often dropped for fault tolerance but mainly for performance. most systems aren't linearizable. it's proven that it's slow because of uncertain delays on the network. thus we go for weaker consistency guarantees to get performance. 

**total order** any two elements can be compared  vs **partial order** some operations are ordered with respect to each other but some are incomparable. 

linearizable systems have total order of operations, whereas causality just defines a partial order. Some operations are ordered relative to each other but others are incomparable. THis is like Git version history -- sometimes commits are one after another, other times they branch and merges are created when commits are combined. 

linearizability implies **causal consistency** (but not the other way around) - causal consistency just means that you know which operations happened before which other operation -- must be able to describe the knowledge of a node in the system (did this know X before Y)

example: collaborative editing - you need to know the order of writes within a document, but not across documents. 

to get causal consistency in a performant way, we need global sequence ordering - logical clock instead of physical clock - example **Lamport timestamp** which is just a tuple of (counter, nodeID), and each request includes the greatest counter read, so if client has read a greater counter from another node, node2 counter is updated to that counter when the request from the client comes in. 

this defines a total order of operations but doesn't give you a way to enforce constraints eg uniqueness constraints. to do this you need **total order broadcast**: protocol for exchanging messages between nodes that ensures no messages are lost and that messages are delivered to all nodes in order (even if node or network is faulty) -- this is implemented by ZooKeeper and etcd. If all messages are delivered and in the same order, machines stay consistent with each other -- this is called **state machine replication**. 

**Two-Phase Commit (2PC)** a way to achieve an atomic transaction across multiple nodes -- all commit or all abort -- via 2 phases: coordinator sends a prepare request - if all reply yes, then send out a commit request, if any say no, abort request is sent. 2 points of "no return" -- when they vote yes, they must be able to commit it later, and once the coordinator decides, it can't go back either and must retry forever until the message gets through. 

Pros and cons of distributed transactions: important for safety, but crappy performance (caused by forcing to disk and more network round trips.) Also if the coordinator dies, locks are being held by the participants causing your application to become unavailable. Basically just best to resolve manually, although there are some ways to add heuristics to recover automatically. 

**Byzantine faults**: when a node in a network deliberately is subverting the system's guarantees. 

**Byzantine Generals Problem**: finding consensus in a network where the nodes may lie or not respond with the truth



### Notes on Raft talk 


