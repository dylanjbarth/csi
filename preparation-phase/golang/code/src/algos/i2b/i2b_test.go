package i2b

import (
	"testing"
)

// Converts a decimal integer to binary representation
func TestI2B(t *testing.T) {
	type testCase struct {
		input  int
		expect []int
	}
	cases := []testCase{
		{1, []int{1}},
		{1, []int{1}},
		{2, []int{1, 0}},
		{3, []int{1, 1}},
		{4, []int{1, 0, 0}},
		{5, []int{1, 0, 1}},
		{6, []int{1, 1, 0}},
		{7, []int{1, 1, 1}},
	}
	for _, c := range cases {
		v := i2b(c.input)
		if len(v) != len(c.expect) {
			t.Errorf("Expected i2b(%d) == %v but got %v", c, c.expect, v)
		}
		for i, cv := range c.expect {
			if v[i] != cv {
				t.Errorf("Expected i2b(%d) == %v but got %v", c, c.expect, v)
			}
		}

	}
}
