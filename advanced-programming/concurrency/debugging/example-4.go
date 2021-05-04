package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{}, 1)
	go func() {
		fmt.Println("performing initialization...")
		done <- struct{}{}
	}()

	<-done
	fmt.Println("initialization done, continuing with rest of program")
}

/*

$ go run -race example-4.go
initialization done, continuing with rest of program
performing initialization...

Problems:

- intended initialization happening _after_ continuation because we are sending to the channel outside the goroutine and receiving inside (so not blocking via the channel as intended)

Solution:

- swap the send and receive on the channel so that we send an event to the done channel when init completes and receive it before continuing execution

$ go run -race example-4.go
performing initialization...
initialization done, continuing with rest of program
*/
