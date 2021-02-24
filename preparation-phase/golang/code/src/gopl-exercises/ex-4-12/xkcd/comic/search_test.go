package comic

import "testing"

func TestSearch(t *testing.T) {
	type testCase struct {
		kw       []string
		i        *SearchIndex
		expected []int
	}
	cases := []testCase{
		{[]string{"barrel"}, &SearchIndex{"barrel": {1: 1}}, []int{1}},
		{[]string{"barrel", "fish"}, &SearchIndex{"barrel": {1: 1}, "fish": {1: 1, 2: 7}}, []int{2, 1}},
		{[]string{"world", "fish"}, &SearchIndex{"barrel": {2: 1}}, []int{}},
	}
	for _, c := range cases {
		out := search(c.kw, c.i)
		if len(out) != len(c.expected) {
			t.Errorf("Expected search(%v) to return slice of len %v but got %d. Results were: %v", c.kw, len(c.expected), len(out), out)
		} else {
			for i, cid := range out {
				if cid != c.expected[i] {
					t.Errorf("Expected search(%v)[%d] to equal %d but got %d.", c.kw, i, c.expected[i], out[i])
				}
			}
		}
	}
}
