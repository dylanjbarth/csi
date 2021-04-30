package exercises

import (
	"fmt"
	"unsafe"
)

// https://research.swtch.com/interfaces
type generic_iface struct {
	tab  *runtime_itab
	data unsafe.Pointer
}

type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      uint8
	align      uint8
	fieldAlign uint8
	kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       int32
	ptrToThis int32
}

type name struct {
	bytes *byte
}

type imethod struct {
	name int32
	ityp int32
}

type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}

type runtime_itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

// Given an interface{} variable that holds an int value, write a function that extracts the int value without using a type assertion or type switch.
func InterfaceExtract(any interface{}) int {
	i := *(*generic_iface)(unsafe.Pointer(&any))
	return *(*int)(i.data)
}

// Given an arbitrary interface value, write a function that iterates through the corresponding itable and prints out information about methods
// You can start by printing just the number of methods, but the eventual goal is for you to explore the underlying representations in an open-ended way.
// We’d recommend time-boxing this exercise because Part 2 will likely address many of your remaining open questions. that can be called on that interface value.
func CountInterfaceMethods(any interface{}) int {
	i := *(*generic_iface)(unsafe.Pointer(&any))
	// fun   [1]uintptr // not too sure how to evaluate this... variable sized array of function pointers and if it's not empty it means it has at least one method.
	// but how do I know how many? maybe try type conversion into a func pointer slice?
	if i.tab.fun[0] == 0 {
		return 0
	}
	// methods := (func)(i.tab.fun[0])
	fmt.Println(i)
	return 5
}

// Now that you know how interfaces are represented in memory, how do you think “type assertions” and “type switches” work? If you were designing Go yourself, how would you approach these features?
/*
	Type assertion: eg someinterface.(someconcretetype) => go probably iterates through the concrete type's itable, and ensures that each method has the same signature.
	Type switches: for each case statement, we are basically just evaluating a type assertion.
*/
