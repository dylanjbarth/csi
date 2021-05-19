package mutex

import (
	"sync/atomic"
)

var nptr uint32

type footex struct {
	locked uint32 // if set, function has called lock and not unlocked yet.
}

// Compare and swap func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
// if *addr == old {
// 	*addr = new
// 	return true
// }
// return false

func (f *footex) Lock() {
	// if locked is nil, the caller gets the lock, otherwise continue to try
	for {
		if atomic.CompareAndSwapUint32(&f.locked, nptr, 1) {
			return
		}
	}
}

func (f *footex) Unlock() {
	if !atomic.CompareAndSwapUint32(&f.locked, 1, nptr) {
		panic("Cannot call unlock on an already unlocked mutex.")
	}
}

// What makes this so slow for 1M??
// lots of contention for the lock, CPU cycles
