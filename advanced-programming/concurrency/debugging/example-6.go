package main

import (
	"fmt"
	"sync"
)

type coordinator struct {
	lock   sync.RWMutex
	leader string
}

func newCoordinator(leader string) *coordinator {
	return &coordinator{
		lock:   sync.RWMutex{},
		leader: leader,
	}
}

// logs state without acquiring a lock -- must acquire lock before calling this!
func (c *coordinator) logStateUnsafe() {
	fmt.Printf("leader = %q\n", c.leader)
}

func (c *coordinator) logState() {
	c.lock.RLock()
	defer c.lock.RUnlock()
	c.logStateUnsafe()
}

func (c *coordinator) setLeader(leader string, shouldLog bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.leader = leader

	if shouldLog {
		c.logStateUnsafe()
	}
}

func main() {
	c := newCoordinator("us-east")
	c.logState()
	c.setLeader("us-west", true)
}

/*
Problem:

$ go run -race example-6.go
leader = "us-east"
^Csignal: interrupt

deadlock. setLeader locks the mutex for reading and writing, and then shouldLog attempts to acquire a lock for reading prior to releasing the outer lock.

Solution:

one way around this is to refactor to decouple lock acquisition from the business logic of logging.

$ go run -race example-6.go
leader = "us-east"
leader = "us-west"
*/
