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

func stringsShareMemory2(s1, s2 string) bool {
	// strings are 2 words each, 1st word is pointer to memory location of underlying byte array...
	// so need to figure out how to access that pointer?
	return &s1 == &s2
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
		next := *(*int)(unsafe.Pointer(uintptr(slicemem.ptr) + scalar*unsafe.Sizeof(i)))
		sum += next
	}
	return sum
}

// Implementation without using an interface, just using the raw bytes!
func sumSlice2(n []int) int {
	slicemem1 := (*int)(unsafe.Pointer(&n))
	slicemem2 := (*int)(unsafe.Pointer(&slicemem1))
	slicemem3 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&n))))
	slicemem4 := *(*int)(unsafe.Pointer((uintptr(*(*uint64)(unsafe.Pointer(&n))))))
	slicemem5 := *(*int)(unsafe.Pointer((*(*uint64)(unsafe.Pointer(&n)))))
	slicemem6 := *(*int)(unsafe.Pointer((uintptr((unsafe.Pointer(&n))))))
	slicemem := unsafe.Pointer(unsafe.Pointer(&n))
	slicelen := unsafe.Pointer((uintptr(unsafe.Pointer(&n)) + unsafe.Sizeof(slicemem)))
	firstEl := *(*int)(slicemem)
	actualLen := *(*int)(slicelen)
	fmt.Println(slicemem, slicelen, firstEl, actualLen, slicemem1, slicemem2, slicemem3, slicemem4, slicemem5, slicemem6)
	// sum := 0
	// for i := 0; i < slicemem.,len; i++ {
	// 	scalar := uintptr(i)
	// 	next := *(*int)(unsafe.Pointer(uintptr(slicemem.ptr) + scalar*unsafe.Sizeof(i)))
	// 	sum += next
	// }
	return 0
}

// local copy of hmap from runtime/map.go
type hmap_copy struct {
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)
}

type mapIface struct {
	_type unsafe.Pointer
	data  unsafe.Pointer
}

func sumMap(m map[int]int) (int, int) {
	var keysum int
	var valsum int
	// thx rishi && https://hackernoon.com/some-insights-on-maps-in-golang-rm5v3ywh
	ei := (*mapIface)(unsafe.Pointer(&m))
	// mdata := (*hmap_copy)(ei.data)
	mtype := (*hmap_copy)(ei._type)
	if mtype.count > 0 {
		nbuckets := mtype.B * mtype.B
		fmt.Println(nbuckets)
		// TODO search through buckets pointer?
	}
	// normally this would be:
	return keysum, valsum
}
