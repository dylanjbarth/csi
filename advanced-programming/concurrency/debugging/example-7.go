package main

import (
	"fmt"
	"sync"
)

const (
	numGoroutines = 100
	numIncrements = 100
)

type counter struct {
	count int
}

func safeIncrement(lock *sync.Mutex, c *counter) {
	lock.Lock()
	defer lock.Unlock()

	c.count += 1
}

func main() {
	var globalLock sync.Mutex
	c := &counter{
		count: 0,
	}

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < numIncrements; j++ {
				safeIncrement(&globalLock, c)
			}
		}()
	}

	wg.Wait()
	fmt.Println(c.count)
}

/*

Problem

$ go run -race example-7.go
==================
WARNING: DATA RACE
Read at 0x00c000138010 by goroutine 8:
  main.safeIncrement()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-7.go:21 +0xbe
  main.main.func1()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-7.go:37 +0xa5

Previous write at 0x00c000138010 by goroutine 7:
  main.safeIncrement()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-7.go:21 +0xd7
  main.main.func1()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-7.go:37 +0xa5

Goroutine 8 (running) created at:
  main.main()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-7.go:33 +0x129

Goroutine 7 (finished) created at:
  main.main()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-7.go:33 +0x129
==================
9654
Found 1 data race(s)
exit status 66

Solution:

- pass by value by default in golang, so we were creating a copy of the lock each time we pass it in.
- turn this into a pointer so we are sharing the same lock between goroutines.

$ go run -race example-7.go
10000
*/
