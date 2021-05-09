package mutex

import (
	"bytes"
	"fmt"
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
// func unsafeGetGoroutineId() uint64 {
// 	b := make([]byte, 64)
// 	b = b[:runtime.Stack(b, false)]
// 	b = bytes.TrimPrefix(b, []byte("goroutine "))
// 	b = b[:bytes.IndexByte(b, ' ')]
// 	n, _ := strconv.ParseUint(string(b), 10, 64)
// 	return n
// }

func (f *footex) Lock() {
	// if lock owner is nil, the caller gets the lock, otherwise continue to try
	pc, _, _, ok := runtime.Caller(1) // get caller ptr
	if !ok {
		panic("Unable to determine ptr to caller of Lock")
	}
	fmt.Printf("Called from %d\n", pc)
	for {
		if atomic.CompareAndSwapUintptr(&f.owner, nptr, pc) {
			fmt.Println("Locked")
			return
		}
	}
}

func (f *footex) Unlock() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("Unable to determine caller ptr")
	}
	fmt.Printf("Called from %d\n", pc)
	for {
		if atomic.CompareAndSwapUintptr(&f.owner, pc, nptr) {
			fmt.Println("Unlocked")
			return
		}
	}
}
