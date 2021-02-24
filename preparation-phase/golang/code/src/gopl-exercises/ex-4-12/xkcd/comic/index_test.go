package comic

import (
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
		input    Comic
		expected []string
	}
	cases := []testCase{
		{Comic{
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
		res := c.input.tokenize()
		if len(res) != len(c.expected) {
			t.Errorf("Expected len(tokenize(%v)) == %d but got %d. Result was %v", c.input, len(c.expected), len(res), res)
		} else {
			for i, tok := range res {
				if tok != c.expected[i] {
					t.Errorf("Expected token at %d to equal %s but got %s", i, c.expected[i], tok)
				}
			}
		}
	}
}

func TestAddToIndex(t *testing.T) {
	type testCase struct {
		input    []Comic
		expected SearchIndex
	}

	cases := []testCase{
		{[]Comic{{
			Month:      "1",
			Num:        1,
			Link:       "",
			Year:       "2006",
			News:       "",
			SafeTitle:  "Barrel - Part 1",
			Transcript: "Barrel barrel barrel",
			Alt:        "Don't we all.",
			Img:        "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
			Title:      "Barrel - Part 1",
			Day:        "1",
		}, {
			Month:      "1",
			Num:        2,
			Link:       "",
			Year:       "2006",
			News:       "",
			SafeTitle:  "Barrel - Part 1",
			Transcript: "Barrel barrel floating ocean.",
			Alt:        "",
			Img:        "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
			Title:      "Barrel - Part 1",
			Day:        "1",
		}}, SearchIndex{"1": {2: 1, 1: 1},
			"2006":     {1: 1, 2: 1},
			"all":      {1: 1},
			"barrel":   {1: 4, 2: 3},
			"dont":     {1: 1},
			"floating": {2: 1},
			"ocean":    {2: 1},
			"part":     {1: 1, 2: 1},
			"we":       {1: 1},
		}}}
	for _, c := range cases {
		idx := make(SearchIndex)
		for _, com := range c.input {
			idx.add(&com)
		}
		if len(idx) != len(c.expected) {
			t.Errorf("Expected len(idx) == %d but got %d. Result was %v", len(c.expected), len(idx), idx)
		} else {
			for k, idxv := range idx {
				if ev, ok := c.expected[k]; !ok {
					t.Errorf("Unexpected token '%s' found in index.", k)
				} else {
					if len(idxv) != len(ev) {
						t.Errorf("Expected len(idx[%s]) == %d but got %d. Result was %v", k, len(c.expected[k]), len(idx[k]), idx[k])
					}
					for kk, idxvv := range idxv {
						if evv, ok := ev[kk]; !ok {
							t.Errorf("Unexpected comic id '%d' found under token '%s' in index.", kk, k)
						} else if evv != idxvv {
							t.Errorf("Expected count of token '%s' for comic id %d to equal %d but got %d.", k, kk, evv, idxvv)
						}
					}
				}
			}
		}
	}
}
