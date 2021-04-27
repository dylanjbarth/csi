package exercises

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

func TestFloat64ToUint64Bin(t *testing.T) {
	n := float64(3.14)
	o := float64ToUint64Bin(n)
	ns := fmt.Sprintf("%064b", math.Float64bits(n))
	os := fmt.Sprintf("%064b", o)
	if ns != os {
		t.Errorf("Expected float64 binary %s to equal uint64 binary %s. But it didn't :(", ns, os)
	}
}

func TestStringsShareMemory(t *testing.T) {
	s1 := "foo"
	s2 := "foo"
	if stringsShareMemory(&s1, &s2) {
		t.Errorf("Expected s1 (%s) and s2 (%s) to not share memory because they are different values and should have different memory locations.", s1, s2)
	}

	s3 := "foo2"
	var s4 *string = (*string)(unsafe.Pointer(&s3)) // todo why does this work?
	// // s4 := *(*string)(unsafe.Pointer(&s3))  // but this doesn't? seems like the one liner de-ref is copying elsewhere in memory?

	if !stringsShareMemory(&s3, s4) {
		t.Errorf("Expected s3 (%s) and s4 (%s) to share memory because s4 is a copy of the memory location at s3.", s3, *s4)
	}

	s5 := "foo3"
	s6 := s5[1:2]
	if !stringsShareMemory(&s5, &s6) {
		t.Errorf("Expected s5 (%s) and s6 (%s) to share memory because s6 is a slice of s5.", s5, s6)
	}
}

func TestSumSlice(t *testing.T) {
	tcase := []int{1, 2}
	total := 3
	out := sumSlice(&tcase)
	if out != total {
		t.Errorf("Expected sum to be %d but got %d", total, out)
	}
}

func TestSumMap(t *testing.T) {
	tcase := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	kt := 3
	vt := 3
	k, v := sumMap(tcase)
	if kt != k {
		t.Errorf("Expected key sum to be %d but got %d. Expected ", kt, k)
	}
	if vt != v {
		t.Errorf("Expected value sum to be %d but got %d. Expected ", vt, v)
	}
}
