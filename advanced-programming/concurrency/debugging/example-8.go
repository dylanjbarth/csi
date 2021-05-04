package main

import (
	"fmt"
	"sync"
	"time"
)

type dbService struct {
	lock       *sync.RWMutex
	connection string
}

func newDbService(connection string) *dbService {
	return &dbService{
		lock:       &sync.RWMutex{},
		connection: connection,
	}
}

// Must establish read lock before calling this!
func (d *dbService) logStateUnsafe() {
	fmt.Printf("connection %q is healthy\n", d.connection)
}

func (d *dbService) takeSnapshot() {
	d.lock.RLock()

	fmt.Printf("Taking snapshot over connection %q\n", d.connection)

	// Simulate slow operation
	time.Sleep(time.Second)

	d.logStateUnsafe()
	d.lock.RUnlock()
}

func (d *dbService) updateConnection(connection string) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.connection = connection
}

func main() {
	d := newDbService("127.0.0.1:3001")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		d.takeSnapshot()
	}()

	// Simulate other DB accesses
	time.Sleep(200 * time.Millisecond)

	wg.Add(1)
	go func() {
		defer wg.Done()

		d.updateConnection("127.0.0.1:8080")
	}()

	wg.Wait()
}

/*
Problem

$ go run -race example-8.go
Taking snapshot over connection "127.0.0.1:3001"
c^Csignal: interrupt

deadlock establishing a read lock between the create snapshot and the logState.. just releasing the lock before reading the state works. Or depending on the original intention, could removing the locking from logstate and just release lock after we've logged state.

*/
