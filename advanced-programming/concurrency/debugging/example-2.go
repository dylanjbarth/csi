package main

import (
	"fmt"
)

const numTasks = 3

func main() {
	done := make(chan struct{})
	for i := 0; i < numTasks; i++ {
		go func() {
			fmt.Println("running task...")

			// Signal that task is done
			done <- struct{}{}
		}()
	}

	// Wait for tasks to complete
	for i := 0; i < numTasks; i++ {
		<-done
	}
	fmt.Printf("all %d tasks done!\n", numTasks)
}

/*
Inital output:

$ go run -race example-2.go
running task...
running task...
running task...
^Csignal: interrupt

never completes...

Solution:
channel was declared but not initialized -- after actually creating the channel it works.

$ go run -race example-2.go
running task...
running task...
running task...
all 3 tasks done!

*/
