package mutex

import (
	"bytes"
	"runtime"
	"strconv"
	"sync/atomic"
)

var nptr uintptr

type footex struct {
	owner uintptr // if set, function has called lock and not unlocked yet.
}

// Compare and swap func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
// if *addr == old {
// 	*addr = new
// 	return true
// }
// return false

// Research on how to get the goroutines ID
// https://pkg.go.dev/github.com/davecheney/junk/id
// If you use this package, you will go straight to hell. LOL
// Pure go version https://blog.sgmansfield.com/2015/12/goroutine-ids/
func unsafeGetGoroutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// fmt.Printf("Stack: %s\n", b)
	// NB this dumps something like:
	/*
		goroutine 8 [running]:
		ap/mutexes.unsafeGetGoroutineId(0x0)
			/Us
		Stack: goroutine 8 [running]:
		ap/mutexes.unsafeGetGoroutineId(0xc00005c
		Stack: goroutine 7 [running]:
		ap/mutexes.unsafeGetGoroutineId(0x0)
			/Us
		Stack: goroutine 7 [running]:
		ap/mutexes.unsafeGetGoroutineId(0xc00005c
	*/
	// And we take the ID of the calling goroutine.
	// Under the hood, runtime.Stack is calling getg => func getg() *g, and formatting this into
	// goroutine + gp.goid + [status]
	// Interesting that goroutines are identity-less from a public perspective but of course they have an ID under the hood!
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func (f *footex) Lock() {
	// if lock owner is nil, the caller gets the lock, otherwise continue to try
	gid := unsafeGetGoroutineId()
	// fmt.Printf("Called from %d\n", gid)
	for {
		if atomic.CompareAndSwapUintptr(&f.owner, nptr, uintptr(gid)) {
			// fmt.Println("Locked")
			return
		}
	}
}

func (f *footex) Unlock() {
	gid := unsafeGetGoroutineId()
	// fmt.Printf("Called from %d\n", gid)
	for {
		if atomic.CompareAndSwapUintptr(&f.owner, uintptr(gid), nptr) {
			// fmt.Println("Unlocked")
			return
		}
	}
}
