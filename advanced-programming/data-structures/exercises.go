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

func sumSlice(n []int) int {
	// get the start of the slice
	startptr := unsafe.Pointer(&n)
	lenptr := unsafe.Pointer(uintptr(startptr) + unsafe.Sizeof(startptr))
	// capptr := unsafe.Pointer(uintptr(startptr) + 2*unsafe.Sizeof(startptr))
	sum := 0
	end := *(*int)(lenptr)
	// a := *(*int)(unsafe.Pointer(&n)) // todo why doesn't this work?? why do I have to get the first index?
	// b := *(*int)(unsafe.Pointer(&n[0]))
	// fmt.Println(a, b)
	for i := 0; i < end; i++ {
		scalar := uintptr(i)
		// sizeof := unsafe.Sizeof(startptr)
		// startmem := uintptr(unsafe.Pointer(&n))
		// finalmem := startmem + scalar*sizeof
		// next := *(*int)(unsafe.Pointer(finalmem))
		next := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&n[0])) + scalar*unsafe.Sizeof(i)))
		sum += next
	}
	return sum
}

func sumMap(m map[int]int) (int, int) {
	return 5, 5
}
