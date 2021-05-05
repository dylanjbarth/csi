package main

import (
	"fmt"
	"sync"
)

var token sync.Mutex

func main() {
	for i := 0; ; {
		token.Lock()
		if i >= 1000 {
			break
		}
		go func() {
			fmt.Printf("launched goroutine %d\n", i)
			i++
			token.Unlock()
		}()
	}
	// Wait for goroutines to finish
	// time.Sleep(time.Second)
}

/*
	Inital Output:
$ go run -race example-1.go
launched goroutine 2
==================
WARNING: DATA RACE
Read at 0x00c0000b8008 by goroutine 7:
  main.main.func1()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-1.go:11 +0x3c

Previous write at 0x00c0000b8008 by main goroutine:
  main.main()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-1.go:9 +0xa4

Goroutine 7 (running) created at:
  main.main()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-1.go:10 +0x7e
==================
launched goroutine 2
launched goroutine 3
launched goroutine 3
launched goroutine 5
launched goroutine 6
launched goroutine 8
launched goroutine 9
launched goroutine 9
launched goroutine 10
Found 1 data race(s)
exit status 66

Problem:
- loop is executing faster than the goroutines are starting so i is incremented faster than individual goroutines can read it.
- time.Sleep isn't a reliable way to ensure that all goroutines finish before the program terminates. If the loop was larger, a second might not be long enough.

Solutions:
- use a sync.Mutex for i
- write i to a "job" channel and create goroutines by reading from that channel
- even simpler -- just pass i to the goroutine...

when using the sync.Mutex approach, we have to move any reads of the shared variables inside the lock. This is ugly but its concurrency safe!
*/
