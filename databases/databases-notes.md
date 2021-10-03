## Key Questions

* what is an RDBMS?
    * type of DBMS - general purpose structured storage and retrieval of data
    * allows us to store and access data in relation to other pieces of data - generally organized into tables with rows and columns. 
    * RDBMS is program for creating and administering the relational database. 
* what problems are they intended to solve?
    * CRUD operations on data 
    * ACID compliance 
    * authentication and authorization 
    * allow you to store data together that has the same attributes (eg a relation / table contains tuples/rows which share the same attributes)
* how are they typically structured?

## architecture of a database system

> This paper presents an architectural discussion of DBMS design principles, including process models, parallel architecture, storage system design, transaction system implementation, query processor and optimizer architectures, and typical shared components and utilities.

DBMS - database management systems

RDBMS - relational database management systems

Questions: 

* why do DBMS systems exist?
    * create an abstraction for application developers so they can store and retrieve data without having to think too hard about the details of how it's persisted and stored. 
* what is the strict definition of an RDBMS vs other types of DBMS?
* what are the most popular RDBMSs in the market? 
    * see https://en.wikipedia.org/wiki/Relational_database#Market_share - Oracle, MySQL, SQL Server, and Postgres.
* what are the 5 major architectural components of an RDBMS?

Components. 

1. Client Communications Manager: job is to establish and persist connection state for the caller, respond to SQL commands, and return data and control messages. client-server interaction, two-tier (direct connection) or three-tier (via proxy). 
2. Process manager: assign thread of computation and get the query scheduled - responsible for admission control. 
3. Relational Query Processor: authorization check and then compiles the query plan, then executes it. 
4. Transactional Storage Manager: manages all data access, enforces ACID compliance. 
5. Shared components and libraries. 


### Process models

- process per worker
- thread per worker 
- thread/process pooling

thrashing is caused by the DBMS not being able to keep the working set of DB pages in memory, and is spending all it's time replacing in-memory pages. can also be caused by lock contention. 

### Parallel architectures 

* shared-memory - all CPUs can access same ram and disk, basically big beefy machine model. 
* shared-nothing - independent machines communicating over network, normally with a single process model per machine, each system will only store a portion of the data. uses horizontal data partitioning to allow nodes to support each other running independently. Requires laying out tables intelligently for efficient sharding. 
* shared-disk - processors can access the disk but not RAM. 
* NUMA - non-uniform memory access - local private memory and shared remote memory. 


### Relational Query Processor 

* DML - data manipulation language - eg SELECT, INSERT, UPDATE, DELETE
* DDL - data definition language - eg CREATE TABLE, CREATE INDEX

DML is processed by the query processor, wheras DDL is implemented procedurally in the DBMS. 

FQDN - server.database.schema.table 

#### Data storage

B+ trees - fast writes, slightly slower reads. 
Bitmap indexes - more efficient storage and potentially faster queries with 

MVCC - multi-version concurrency control: transactions are completed in order, each transaction can only see results of transactions completed before it started. 


## Class 1 Notes

lldb 
attach

flex, bison: generic DSL parsers (used by postgres to parse the text into a syntax tree)

The lifecycle of a query: https://www.postgresql.org/docs/devel/query-path.html, https://www.postgresql.org/developer/backend/, 

pg internals https://momjian.us/main/writings/pgsql/internalpics.pdf


## Sorting and Hashing Lecture 

### why sorting 
- sort to eliminate duplicates, summarize groups, or just because it's been requested in order. 
- sort-merge join algorithm, sort before creating an index

how can you sort 100GB with just 1GB of RAM? relying on OS / virtual memory would lead to page faulting lots of random IO, have to be strategic about IO, minimize number of calls made. 

performance changes depending on the hardware -- magnetic disks vs flash -- 10x cheaper. Sometimes flash is used as an intermediary cache, eventually it will probably just become the main persistence medium. 

Magnetic disk is literally spinning platters with an arm reading/writing/seeking. Sectors on each disk like a pizza slice, this is a block/page size. Seek movement (arm movement). Time to access a disk block: seek time = 2-4ms, rotational delay (waiting for block to rotate to the arm), another 2-4ms, and then transfer time is .3ms per 64kb. 

To reduce IO time - try and arrange pages on disk so that they are on the same track/cylinder etc. 

On Flash, random and sequential reads are about the same and fast, random writes are slower than sequential writes. Point is that there is still a cost for randomness. 

Sorting: 

2-Way sorting, conquer and merge -- read page, sort it, write it. Then repeat for all pages. In the second pass, merge two sorted blocks at a time. Then repeat. 

External merge sort: >3 buffer pages, log base b-1(N/B). 

Memory requirement for external sort To sort N pages you need sqrt(N) space to just take 2 passes through the data. 

Heapsort aka tournament sort. 

Hashing when we don't require order, just require groups or removal of duplicates, just want to rendevous matches. 

partition the dataset into buckets using a course hash function, then read each partition into RAM hash table individually. Run into issues when group size is bigger than your partition. 

hashing is divide and conquer and sorting is conquer and merge. 

## Single Table Queries lecture 

Relational tables -- give it a schema which is fixed, attribute names, atomic types. Records can change (rows, tuples) stored as multisets. 

basic single table query in SQL 

```SQL
SELECT columns FROM table WHERE predicate GROUP BY columns HAVING predicate ORDER BY columns
```

Query Executor Architecture 

1. Query Optimizer => Query Plan => Relational Operators => Executor 

Query optimizer translates SQL to query plans - query executor is interpreter for query plans - query plan is like blobs and arrows - dataflow diagrams - each node is a relational operator with tuples flowing between them. How data is flowing through chunks of code. 

SELECT DISTINCT name, gpa from STUDENTS turns into: 

1. File scan => name, gpa tuples
2. Sort => name, gpa
3. Distinct => name, gpa 

each **relational operator** is a subclass of an Iterator class (init, next(), close()) and any iterator's output can be input to another iterator, so they can all connect to each other. 

Aggregate - maintains a per group depending on the relational operator 
* Count
* Sum 
* Average (keeps sum and count)

Sorted group by vs a hash group by 

- only difference is that the batches aren't in order, but like items are grouped together. 

Relational operators: 

- Projection: select certain columns from tuples
- Selection: filter tuples based on a predicate
- Scan: 
- Limit: 

Class Notes

- 30k rows, sort by genre, what sorting algorithm with postgres use: 
- quicksort: favored if the dataset can fit into memory. 
- heapsort: sometimes fatser than quicksort, but can be slower in the worst case. When you combine a limit with a sort, heapsort can be much more memory efficient because you can maintain the heap and discard anything beyond the limit. 
- external mergesort: favored if the dataset will not all fit into memory

Explain analyze: 
- cost is only useful in relative terms, there are two numbers. Cost to get first row and cost to get all the rows. Time is estimated in ms. 

# Physical Storage
 
how does the database interact with the file system? how does it manage disk access and memory buffers? 
what is the "heap file" format and how is it used to store unordered records?

why not just use the OS file system as a database? 
- OS doesn't help with concurrent access, undeterministic behavior
- OS doesn't guarantee durability on crash or power outage (fault tolerance and recovery)

DBMS offers these guarantees, provides abstractions so we can not worry about durability or transaction management. 

Basic patterns for big data: 
- streaming (unix pipes work this way too, OS handles the buffering for you)
- divide and conquer 

"disorder is a friend of scaling" because it allows for reorganizing depending on the purposes, eg optimize for cache locality, work in batches, etc. 

rendevous - need certain items to be co-resident in memory (not guaranteed to appear in same input chunk) - 

"out-of-core" algorithms - algos that operate on data sets that don't fit into memory

divide and conquer -- generally 2 phases, 1) clean up the data on the way in and get it into chunks you want, then 2) conquer on the second pass. 

## Storing Data: Disks and Files 

disk management and buffer management

disk space manager which is managing IO at the lowest level, has access to actual blocks/pages on disk. 

exposes a "file" abstraction, so a file is a collection of pages/blocks containing collection of records, diff from Unix files in that those are sequential bytes, these aren't. 

Supports: insert/delete/modify/fetch/scan 

Implemented as multiple OS files or just raw disk space.  

Heap file: unordered collection of records, have to keep track of pages, free space, and where the records are on the page. 

Can be implemented in two ways: 
- implemented as linked list: database catalog stores the header page which points to page ID and heap file.  
- better way is to use a page directory in header page - keep free space count. 


How to store individual records/tuples in a file: 

- fixed length => just base + offset
- variable length:
    - fields delimited by special symbols (eg end each field with a semicolon)
    - array of field offsets at the header telling you where each field starts and ends. this allows for efficient storage of nulls (2 pointers pointing to the same place)
    
    
How to store pages?

- with fixed length records, just store number of records at end of page so it's easy to figure out where your records stop and free space begins. You can either compact free space when things are removed or not. In cases where you don't compact, you need a bitmap or something to know where the free space is. Requires using slots instead of packing bytes sequentially. Advantage of non-compact strategy is that the record ID stays the same (where it lives on disk). When you compact you have to change the record tuples for everything on the page -- this is a PITA because you have to fix all pointers, eg update indexes. 
- with variable length records => **slotted page format**. pointer at the end showing start of free space. N slots following from end of page, which provide byte offset and length in the slot directory. 

System catalogs (internal tables basically): how to find stuff -- 
- for each table, contains the name, file location, and file structure, attributes, indexes, constraints
- for each index, says the structure and the key fields. 

Views are stored queries that can be addressed like a table. 

Buffer management: caching layer between the disk space manager and then the higher level file API 

basically an array of memory containing frame ID (location in memory table), page ID - what page is in what frame, pin_count - how many tasks have pinned page in memory, dirty (meaning it must be written to disk because it was changed since it was read) 
- what does it mean to pin a page in memory?

only unpinned pages can get replaced. these are normally very short lived. 

DB can predict what pages you need and get pre-fetched (eg during sequential scanning)

page requestor must unpin and flip the dirty bit if it was changed. 

how does buffer manager choose which frame to replace?
- least recently used
    - track time each frame was last unpinned (end of use)
    - replace frame with earliest unpinned time. 
    - this works well for repeated access to popular pages.
    - doesn't work well if you are doing repeated sequential scans -- best you can do here is B-1 so used most recently used. 
- most recently used, 
    - works best in sequential scanning pattern. 
- clock replacement policy
    - provides constant time LRU replacement policy
    - basically we evict the first item in the buffer we've seen twice that has pin count 0. 

replacement policy can have big impact on I/O of system, really depends on the access pattern of your workload. 

why does pinning happen? assumption here is that this is going to be related to transaction management / concurrency control. 

why not use the OS file system? because we want to control the replacement policy, prefetch, pin pages. 

question: is the DBMS using system calls still to access the disk? it must? -- says done via "lowel-level" OS interfaces to avoid the OS file cache and control when we write. 