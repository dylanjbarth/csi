package main

import (
	"fmt"
	"testing"
)

func TestAwesomeBloomFilterHasNoFalseNegatives(t *testing.T) {
	b := newAwesomeBloomFilter(64*4, 1)
	cases := []string{"hey", "this", "is", "blooming", "awesome"}
	for _, c := range cases {
		b.add(c)
		fmt.Printf("%b\n", b.data)
	}
	for _, c := range cases {
		if ok := b.maybeContains(c); !ok {
			t.Errorf("Expected set to contain %s but got %t", c, ok)
		}
	}
}
