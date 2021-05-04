package concurrency

import "sync/atomic"

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
	atomic.AddUint64(&n.counter, 1)
	return n.counter
}
