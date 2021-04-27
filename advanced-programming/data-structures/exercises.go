package exercises

import (
	"fmt"
	"unsafe"
)

func float64ToUint64Bin(n float64) uint64 {
	// 1. Create pointer of arbitrary type using unsafe.Pointer
	// 2. we then do a type conversion, converting the unsafe.Pointer to a uint64 pointer.
	// 3. We then dereference the pointer to return the value which is the same bits in memory as the float64 because the pointers are pointing at the same memory value.
	return *(*uint64)(unsafe.Pointer(&n))
}

type strIface struct {
	ptr unsafe.Pointer
	len int
}

func stringsShareMemory(s1, s2 *string) bool {
	// strings are 2 words each, 1st word is pointer to memory location of underlying byte array...
	// so need to figure out how to access that pointer?
	s1p := *(*strIface)(unsafe.Pointer(s1))
	s2p := *(*strIface)(unsafe.Pointer(s2))
	s1start := uintptr(s1p.ptr)
	s2start := uintptr(s2p.ptr)
	s1end := s1start + uintptr(s1p.len)
	s2end := s2start + uintptr(s2p.len)
	return (s1start <= s2start && s1end >= s2start) || (s1start <= s2end && s1end >= s2end)
}

type sliceIface struct {
	// https://research.swtch.com/godata
	ptr unsafe.Pointer
	len int
	cap int
}

func sumSlice(n []int) int {
	slicemem := (*sliceIface)(unsafe.Pointer(&n))
	sum := 0
	for i := 0; i < slicemem.len; i++ {
		scalar := uintptr(i)
		next := *(*int)(unsafe.Pointer(uintptr(slicemem.ptr) + scalar*unsafe.Sizeof(slicemem.ptr)))
		sum += next
	}
	return sum
}

func sumMap(m map[int]int) (int, int) {
	return 5, 5
}
