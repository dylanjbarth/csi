# Key value store

## Usage

`make server` to start the backend (which spins up a routing tier, one leader, and one follower) && `make client` to start a shell to interact with it. Allows multiple client connections.

## Tests

`make test`

## TODO / Bugs
- checkpoint the file to disk but keep in memory so it doesn't have to be loaded into memory all the time. Create max buffer size to hold in memory and then flush to disk. 
- figure out a use case.
- improve shell UX (arrow keys, meta chars)
- improve parsing, can't set multi word values for example. eg `set foo=bar baz` fails
- how to synchronize in situation where a new node is added or the follower is behind.