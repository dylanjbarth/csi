package index

import (
	"testing"
)

type testCase struct {
	input, expected string
}

func TestClean(t *testing.T) {
	cases := []testCase{
		{"punc.", "punc"},
		{"$p%u%n#c))).", "punc"},
		{"UPPER", "upper"},
		{"123", "123"},
	}
	for _, c := range cases {
		if clean(c.input) != c.expected {
			t.Errorf("Expected clean(%s) == %s but got %s", c.input, c.expected, clean(c.input))
		}
	}
}
