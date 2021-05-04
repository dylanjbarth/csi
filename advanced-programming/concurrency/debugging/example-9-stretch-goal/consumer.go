package main

import (
	"fmt"
	"sync"
)

type Consumer struct {
	id   int
	s    *StateManager
	lock *sync.RWMutex
}

func NewConsumer(id int, s *StateManager) *Consumer {
	return &Consumer{
		id:   id,
		s:    s,
		lock: &sync.RWMutex{},
	}
}

func (c *Consumer) GetState() string {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return fmt.Sprintf("<GetState result for consumer %d>", c.id)
}

func (c *Consumer) Terminate() {
	c.lock.Lock()
	defer c.lock.Unlock()

	// You can imagine that this internal cleanup mutates the state
	// of the Consumer
	fmt.Printf("Performing internal cleanup for consumer %d\n", c.id)

	c.s.RemoveConsumer(c.id)
}
