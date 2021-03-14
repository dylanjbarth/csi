// Say I wanted to calculate the sum of the first n numbers, and I’m wondering how long this will take. Firstly, can you think of a simple algorithm to do the calculation? It should be a function that has n as a parameter, and returns the sum of the first n numbers. You don’t need to do anything fancy, but please do take the time to write out an algorithm and think about how long it will take to run on small or large inputs.

package sum

import "testing"

func TestSum(t *testing.T) {
	type testCase struct {
		input    int
		expected int
	}
	cases := []testCase{{0, 0}, {1, 1}, {2, 3}, {3, 6}, {10, 55}}
	for _, c := range cases {
		v := sum(c.input)
		if v != c.expected {
			t.Errorf("Expected sum(%d) == %d but received %d", c.input, c.expected, v)
		}
	}
}
