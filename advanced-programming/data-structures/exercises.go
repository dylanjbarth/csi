package exercises

import (
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
	// s1ptr := unsafe.Pointer(s1) // this is a pointer to the start of the 2 word string in memory (which itself is a pointer to the underlying data)
	// // s1len := unsafe.Pointer(uintptr(s1ptr) + unsafe.Sizeof(s1ptr))
	// s2ptr := unsafe.Pointer(s2)
	// return s1ptr == s2ptr
	// s2len := unsafe.Pointer(uintptr(s2ptr) + unsafe.Sizeof(s2ptr))
	// fmt.Println(*(*string)(s1ptr))
	// fmt.Println(*(*int)(s1len))
	// s2ptr := unsafe.Pointer(&s2)
	// s1ptr := unsafe.Pointer(&s1)
	// s2ptr := unsafe.Pointer(&s2)
	// return s1ptr == s1len
	// but this still works too?
	return s1 == s2
	// TODO think the actual point of this exercise is to see if any of the memory overlaps, not just the start of the string data!
}

func sumSlice(n *[]int) int {
	// strategy here is to
	return 5
}

func sumMap(m map[int]int) (int, int) {
	return 5, 5
}
