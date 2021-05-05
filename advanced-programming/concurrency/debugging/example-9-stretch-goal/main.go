package main

import (
	"sync"
)

func main() {
	s := NewStateManager(10)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		c := s.GetConsumer(0)
		c.Terminate()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		s.PrintState()
	}()

	wg.Wait()
}

/*
Problem:

$ go run -race main.go consumer.go state_manager.go
Performing internal cleanup for consumer 0
Started PrintState
<GetState result for consumer 2>
<GetState result for consumer 4>
<GetState result for consumer 5>
<GetState result for consumer 8>
<GetState result for consumer 7>
<GetState result for consumer 9>
^Csignal: interrupt

hangs.

*/
