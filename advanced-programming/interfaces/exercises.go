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

type runtime_itab struct {
	// odd that this and the _type were not found by the compiler in the runtime. Where are they?
	// 	inter *interfacetype
	// _type *_type
	// using unsafe pointer for now
	inter unsafe.Pointer
	_type unsafe.Pointer
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
	if (i.tab.fun[0] == 0) {
		return 0
	}
	methods := (func)(i.tab.fun[0])
	fmt.Println(i)
	return 5
}
