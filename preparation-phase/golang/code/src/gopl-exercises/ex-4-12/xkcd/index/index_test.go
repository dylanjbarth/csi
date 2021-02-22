package index

import (
	"gopl-exercises/ex-4-12/xkcd/types"
	"testing"
)

func TestClean(t *testing.T) {
	type testCase struct {
		input, expected string
	}
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

func TestTokenize(t *testing.T) {
	type testCase struct {
		input    types.Comic
		expected []string
	}
	cases := []testCase{
		{types.Comic{
			Month:      "1",
			Num:        1,
			Link:       "",
			Year:       "2006",
			News:       "",
			SafeTitle:  "Barrel - Part 1",
			Transcript: "[[A boy sits in a barrel which is floating in an ocean.]]\nBoy: I wonder where I'll float next?\n[[The barrel drifts into the distance. Nothing else can be seen.]]\n{{Alt: Don't we all.}}",
			Alt:        "Don't we all.",
			Img:        "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
			Title:      "Barrel - Part 1",
			Day:        "1",
		}, []string{"barrel", "part", "1", "a", "boy", "sits", "in", "a", "barrel", "which", "is", "floating", "in", "an", "ocean", "boy", "i", "wonder", "where", "ill", "float", "next", "the", "barrel", "drifts", "into", "the", "distance", "nothing", "else", "can", "be", "seen", "alt", "dont", "we", "all", "dont", "we", "all", "2006"},
		}}
	for _, c := range cases {
		res := tokenize(&c.input)
		if len(res) != len(c.expected) {
			t.Errorf("Expected len(tokenize(%v)) == %d but got %d. Result was %v", c.input, len(c.expected), len(res), res)
		}
		for i, tok := range res {
			if tok != c.expected[i] {
				t.Errorf("Expected token at %d to equal %s but got %s", i, c.expected[i], tok)
			}
		}
	}
}
