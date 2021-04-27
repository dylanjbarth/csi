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

// TODO is there a way this would ever return true if we didn't pass string pointers here due to pass by value nature of go?
func areStringsAliases(s1, s2 *string) bool {
	// strings are 2 words each, 1st word is pointer to memory location of underlying byte array...
	// so need to figure out how to access that pointer?
	return s1 == s2
}

func sumSlice(n []int) int {
	return 5
}

func sumMap(m map[int]int) (int, int) {
	return 5, 5
}
