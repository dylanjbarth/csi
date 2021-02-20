package main

import (
	"testing"
)

func TestReverseP(t *testing.T) {
	rev := []int{3, 2, 1}
	x := []int{1, 2, 3}
	reversep(&x)
	for i, v := range x {
		if v != rev[i] {
			t.Errorf("Expected x[%d] == %d but got %d", i, rev[i], v)
		}
	}
}
