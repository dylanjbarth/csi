package main

import (
	"fmt"
	"math/rand"
	"time"
)

var responses = []string{
	"200 OK",
	"402 Payment Required",
	"418 I'm a teapot",
}

func randomDelay(maxMillis int) time.Duration {
	return time.Duration(rand.Intn(maxMillis)) * time.Millisecond
}

func query(endpoint string) string {
	// Simulate querying the given endpoint
	delay := randomDelay(100)
	time.Sleep(delay)

	i := rand.Intn(len(responses))
	return responses[i]
}

// Query each of the mirrors in parallel and return the first
// response (this approach increases the amount of traffic but
// significantly improves "tail latency")
func parallelQuery(endpoints []string) string {
	results := make(chan string)
	for i := range endpoints {
		go func(e string) {
			results <- query(e)
		}(endpoints[i])
	}
	return <-results
}

func main() {
	var endpoints = []string{
		"https://fakeurl.com/endpoint",
		"https://mirror1.com/endpoint",
		"https://mirror2.com/endpoint",
	}

	// Simulate long-running server process that makes continuous queries
	for i := 0; i < 100; i++ {
		fmt.Println(parallelQuery(endpoints))
		delay := randomDelay(100)
		time.Sleep(delay)
	}
}

/*

$ go run -race example-5.go
418 I'm a teapot
418 I'm a teapot
200 OK
200 OK
200 OK
200 OK
200 OK
200 OK
418 I'm a teapot
418 I'm a teapot

Problem:

- from the code we'd expect results to be evenly distributed, but they don't seem to be...
- race detector isn't detecting anything, but if we look for shared variables inside goroutines, we see that i (the endpoint index) is used inside the goroutine, so that's probably a hint.


Potential solution:

- not 100% sure this is it, but parameterizing the endpoint itself in the goroutine instead of re-using i seems to provide better distribution

$ go run -race example-5.go
418 I'm a teapot
418 I'm a teapot
200 OK
200 OK
402 Payment Required
200 OK
200 OK
200 OK
418 I'm a teapot
418 I'm a teapot

^^^^ this was way off -- memory leak because the other two goroutines are sending and the nothing is receiving on the channel.

*/
