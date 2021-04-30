package exercises

import (
	"fmt"
	"testing"
)

func TestInterfaceExtract(t *testing.T) {
	tcase := []int{1, 2, 4, 10}
	for _, v := range tcase {
		out := InterfaceExtract(interface{}(int(v)))
		if out != v {
			t.Errorf("Expected interface extract to return %d but got %d", out, v)
		}
	}
}

type empty interface{}
type oneM interface {
	one(int) int
}
type twoM interface {
	one(int) int
	two(int) int
}
type thing int

func (t thing) one(a int) int {
	return 1 + a
}
func (t thing) two(a int) int {
	return 2 + a
}
func TestCountInterfaceMethods(t *testing.T) {
	type tcase struct {
		label string
		any   interface{}
		count int
	}
	a := thing(1)
	cases := []tcase{
		{
			"empty",
			empty(1),
			1,
		},
		{
			"oneM",
			oneM(a),
			1,
		},
		{
			"twoM",
			twoM(a),
			2,
		},
	}
	for _, v := range cases {
		fmt.Printf("Running %s", v.label)
		o := CountInterfaceMethods(v.any)
		// ensure the compiler doesn't optimize away the methods?
		switch a := v.any.(type) {
		case twoM:
			fmt.Println(a.one(1))
			fmt.Println(a.two(1))
		case oneM:
			fmt.Println(a.one(1))
		}
		if o != v.count {
			t.Errorf("Expected interface extract to return %d but got %d", v.count, o)
		}
	}
}
