# Key value store

## Usage

`make server` to start the backend && `make client` to start a shell to interact with it. Allows multiple client connections.

## Tests

`make test`

## TODO / Bugs
- checkpoint the file to disk but keep in memory so it doesn't have to be loaded into memory all the time. Create max buffer size to hold in memory and then flush to disk. 
- figure out a use case.
- improve shell UX (arrow keys, meta chars)
- improve parsing, can't set multi word values for example. eg `set foo=bar baz` fails

## Replication design idea: 
- introduce a comms manager (load balancer) which is a thin layer that routes to leader or follower pool based on write/read. This way client only has one address for server and number of followers can scale up or down (just need to register with comms manager and leader)
- after write, leader broadcasts write to all followers. 
observation: complexity really immediately increases as soon as you try and add more than 1 server - introducing a load balancer, figuring out the communication scheme, thinking about how the storage format will scale up or down. 
