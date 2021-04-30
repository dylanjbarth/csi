package exercises

import "testing"

func TestInterfaceExtract(t *testing.T) {
	tcase := []int{1, 2, 4, 10}
	for _, v := range tcase {
		out := InterfaceExtract(interface{}(int(v)))
		if out != v {
			t.Errorf("Expected interface extract to return %d but got %d", out, v)
		}
	}
}
