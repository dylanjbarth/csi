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
	// fmt.Printf("Called from %d\n", gid)
	for {
		if atomic.CompareAndSwapUint32(&f.locked, nptr, 1) {
			// fmt.Println("Locked")
			return
		}
	}
}

func (f *footex) Unlock() {
	// fmt.Printf("Called from %d\n", gid)
	for {
		if atomic.CompareAndSwapUint32(&f.locked, 1, nptr) {
			// fmt.Println("Unlocked")
			return
		}
		panic("Cannot call unlock on an already unlocked mutex.")
	}
}
