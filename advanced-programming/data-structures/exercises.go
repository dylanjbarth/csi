package exercises

import "unsafe"

func float64ToUint64Bin(n float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&n))
}

func areStringsAliases(s1, s2 string) bool {
	// strings are 2 words each, 1st word is pointer to memory location of underlying byte array...
	// so need to figure out how to access that pointer?
	return &s1 == &s2
}

func sumSlice(n []int) int {
	return 5
}

func sumMap(m map[int]int) (int, int) {
	return 5, 5
}
