# Notes 

- system design requires background knowledge as opposed. 
- example "design youtube" and leads to 45 minute discussion by prodding interviewer
- unlike coding interviews, lot of subjectivity in system design. not objectively correct or incorrect 
- confidently justify and rationalize choices and tradeoffs between systems. 

# Client server architecture 
- client: makes requests for data 
- server: responds to requests for data
- HTTP request to ports 

"how does the internet work" - client under the hood makes a DNS request to figure out the IP address.  

# Network protocols 

- protocols: agreed upon communication process. types of messages, order, etc. 
- IP, TCP, HTTP: IP = how you get there, TCP = in-order packet deliver, HTTP = application layer business logic data passing. 

packet IP header: destination and sender address, total size, version of IP protocol, etc. ipv4 and ipv6 - 32 bytes vs 128 bytes

# Storage 
- difference between desk and memory... 
- availability vs structure of data etc. 

# Latency & Throughput
most important measures of performance
- latency: how long does it take for data to get from point A to point B in the system. (eg network request from client to server and response, or reading from disk). think orders of magnitude: 1MB from memory == 250 microseconds. 1MB from SSD == 1000 microseconds. 1MB over 1Gbps network == 10,000 microseconds. 1MB on HDD == 20,000 microseconds. packet from CA => Neth => CA == 150,000 microseconds. 
  - Things that need low latency: video games, trading platforms, etc. 
- throughput: how much work can a machine perfom in a given amount of time (how much data can it transfer from point A to point B in a given amount of time) - measured in bytes per second. 

throughput and latency aren't necessarily correlated, so don't assume throughput from latency and vice versa, they are distinct. 

# Availability 
- how resistant is the system to failures? fault tolerance. in a given period of time, what % of the time is it up? 
- ask: what is the important of fault tolerance for the system requirements? eg airplane software vs Youtube vs cloud providers vs a personal blog. 
- "nines" - 5 nines = 5 minutes of downtime per year (gold standard = highly available), 4 nines = ~ 1 hour of downtime per year. 
- SLA = guarantee of uptime, SLO = component of SLA, the guaranteed % of uptime as an example. 
- how do you improve availability? identify and minimize single points of failure through redundancy. 

# Caching 
- something that comes up in almost every single systems design interview
- avoid redoing operations and thus speed up a system by reducing latency
- definition: storing data in a way that makes it faster to access to reduce latency. 
- example: avoid making a network request by caching on the client side or server side to skip the database hit. 

caching in action: 
- write-through cache: when you write it, write to cache and persist in same operation to keep the cache fresh. 
- write-back cache: server updates cache, then async update the database with the values in the cache. 

caches become stale if they haven't been updated properly: 
- one solution is to condense cache to one location (eg comments that can be edited)
- ask yourself: does staleness matter? is the data mutable or immutable? 

eviction policies: 
- how do you get rid of data from the cache? 
- least recently used - LRU 
- least frequently used - LFU 
- LIFO, FIFO 
- talk with interviewer to figure out what the requirement is

# Proxies
- forward proxy: server sitting between client and server - acts on behalf of the clients (on the client team) - client hits forward proxy to hit the server. 
  - basically how VPN works
- reverse proxy: act on behalf of server, client unknowingly hits the reverse proxy which forwards to the server. 
  - useful for: filtering requests, do logging, cache stuff, or load balancing 

# Load balancing 
- server that sits between client and server that is responsible for balancing the workload between servers. 
- hardware (limited to what you are given) vs software (more configuration) load balancers 
- servers register themselves with the load balancer 
- load balancing strategies: 
  - randomly
  - round-robin: goes through in a loop, guarantee even distribution of traffic
  - weighted round-robin: weight specific servers, useful in cases where one server is more powerful than others 
  - performance-based: LB runs health checks on the server and redirects accordingly
  - client IP based - hash the client IP and direct to specific server (useful to maintain server caches ensuring client reaches the same server each time)

# Hashing 
- transform data into a fixed size value. issues when you add servers, because you increase cache misses 
- consistent hashing:  organize the servers into a circular buffer (put them on a clock) - hash the servers onto the buffer. Do the same thing with your clients. This means when you add/remove clients or servers you minimize issues -- client gets the server closest in a clockwise direction. 
- rendevous hashing: rank/score servers for clients, if a server disappears then use the next highest scoring one (could be a hashing function or a scoring function)
in both cases, you reduce remapping churn when servers are added or removed. 

# Relational databases
- relational = tabular, structures for storing data representing an entity 