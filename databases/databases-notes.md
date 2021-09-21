## architecture of a database system

> This paper presents an architectural discussion of DBMS design principles, including process models, parallel architecture, storage system design, transaction system implementation, query processor and optimizer architectures, and typical shared components and utilities.

DBMS - database management systems

RDBMS - relational database management systems

Questions: 

* why do DBMS systems exist?
    * create an abstraction for application developers so they can store and retrieve data without having to think too hard about the details of how it's persisted and stored. 
* what is the strict definition of an RDBMS vs other types of DBMS?
* what are the most popular RDBMSs in the market? 
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