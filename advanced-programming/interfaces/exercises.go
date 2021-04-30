package exercises

import (
	"unsafe"
)

// https://research.swtch.com/interfaces
type generic_iface struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}

// Given an interface{} variable that holds an int value, write a function that extracts the int value without using a type assertion or type switch.
func InterfaceExtract(any interface{}) int {
	i := *(*generic_iface)(unsafe.Pointer(&any))
	return *(*int)(i.data)
}
