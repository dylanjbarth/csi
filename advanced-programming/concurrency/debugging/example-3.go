package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var urls = []string{
		"https://bradfieldcs.com/courses/architecture/",
		"https://bradfieldcs.com/courses/networking/",
		"https://bradfieldcs.com/courses/databases/",
	}
	var wg sync.WaitGroup
	for i := range urls {
		wg.Add(1)
		go func(i int) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			_, err := http.Get(urls[i])
			if err != nil {
				panic(err)
			}

			fmt.Println("Successfully fetched", urls[i])
		}(i)
	}

	// Wait for all url fetches
	wg.Wait()
	fmt.Println("all url fetches done!")
}

/*
Initial output:

$ go run -race example-3.go
==================
WARNING: DATA RACE
Write at 0x00c000138288 by main goroutine:
  internal/race.Write()
      /usr/local/go/src/internal/race/race.go:41 +0x125
  sync.(*WaitGroup).Wait()
      /usr/local/go/src/sync/waitgroup.go:128 +0x126
  main.main()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-3.go:32 +0x18e

Previous read at 0x00c000138288 by goroutine 7:
  internal/race.Read()
      /usr/local/go/src/internal/race/race.go:37 +0x206
  sync.(*WaitGroup).Add()
      /usr/local/go/src/sync/waitgroup.go:71 +0x219
  main.main.func1()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-3.go:18 +0x5e

Goroutine 7 (running) created at:
  main.main()
      /Users/dylanbarth/projects/csi/advanced-programming/concurrency/debugging/example-3.go:17 +0x170
==================
Successfully fetched https://bradfieldcs.com/courses/databases/
Successfully fetched https://bradfieldcs.com/courses/architecture/
Successfully fetched https://bradfieldcs.com/courses/networking/
all url fetches done!
Found 1 data race(s)
exit status 66

Problems:

- we are incrementing the wg inside the goroutine (async) instead of before we start the goroutine, causing a race condition.

Solution:

- move wg.add call outside of the goroutine (but inside the loop)

$ go run -race example-3.go
Successfully fetched https://bradfieldcs.com/courses/architecture/
Successfully fetched https://bradfieldcs.com/courses/databases/
Successfully fetched https://bradfieldcs.com/courses/networking/
all url fetches done!
*/
