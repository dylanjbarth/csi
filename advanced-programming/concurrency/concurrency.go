package concurrency

import (
	"sync"
	"sync/atomic"
)

type counterService interface {
	// Returns values in ascending order; it should be safe to call
	// getNext() concurrently without any additional synchronization.
	getNext() uint64
}

type NoSync struct {
	counter uint64
}

func (n *NoSync) getNext() uint64 {
	n.counter += 1
	return n.counter
}

type AtomicCount struct {
	counter uint64
}

func (n *AtomicCount) getNext() uint64 {
	result := atomic.AddUint64(&n.counter, 1)
	return result
}

type SyncMutexCount struct {
	token   sync.Mutex
	counter uint64
}

func (n *SyncMutexCount) getNext() uint64 {
	n.token.Lock()
	defer n.token.Unlock()
	n.counter += 1
	return n.counter
}

type ChannelCount struct {
	requests  chan struct{}
	responses chan uint64
}

func initChannelCounter(start uint64) *ChannelCount {
	cc := new(ChannelCount)
	cc.requests = make(chan struct{})
	cc.responses = make(chan uint64)
	go func() {
		var count = start
		for {
			<-cc.requests
			count++
			cc.responses <- count
		}
	}()
	return cc
}

func (cc *ChannelCount) getNext() uint64 {
	cc.requests <- struct{}{}
	return <-cc.responses
}
