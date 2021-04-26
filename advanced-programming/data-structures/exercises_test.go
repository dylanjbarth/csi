package exercises

import (
	"fmt"
	"math"
	"testing"
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

func TestAreStringsAliases(t *testing.T) {
	s1 := "foo"
	s2 := s1
	if !areStringsAliases(s1, s2) {
		t.Errorf("Expected s1 (%d) and s2 (%d) to be aliases because they have the same memory location.", &s1, &s2)
	}
}

func TestSumSlice(t *testing.T) {
	tcase := []int{1, 2}
	total := 3
	out := sumSlice(tcase)
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
