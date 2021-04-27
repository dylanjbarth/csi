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

func stringsShareMemory(s1, s2 *string) bool {
	// strings are 2 words each, 1st word is pointer to memory location of underlying byte array...
	// so need to figure out how to access that pointer?
	s1ptr := unsafe.Pointer(s1) // this is a pointer to the start of the 2 word string in memory (which itself is a pointer to the underlying data)
	s1len := unsafe.Pointer(uintptr(s1ptr) + unsafe.Sizeof(s1ptr))
	s1start := uintptr(s1ptr)
	s1end := uintptr(s1ptr) + uintptr(s1len)*unsafe.Sizeof(s1ptr)
	s2ptr := unsafe.Pointer(s2)
	s2len := unsafe.Pointer(uintptr(s2ptr) + unsafe.Sizeof(s2ptr))
	s2start := uintptr(s2ptr)
	s2end := uintptr(s2ptr) + uintptr(s2len)*unsafe.Sizeof(s2ptr)
	fmt.Printf("Test case s1 %s start %d end %d s2 %s start %d end %d\n", *s1, s1start, s1end, *s2, s2start, s2end)
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
