package tempconv

import "testing"

func TestKToC(t *testing.T) {
	exp := Celsius(-268.15)
	i := Kelvin(5.)
	c := KToC(i)
	if c != exp {
		t.Errorf(`KToC(%g) should return %g but got %g`, i, exp, c)
	}
}

func TestKToF(t *testing.T) {
	exp := Fahrenheit(-450.66999999999996)
	i := Kelvin(5.)
	c := KToF(i)
	if c != exp {
		t.Errorf(`KToF(%g) should return %g but got %g`, i, exp, c)
	}
}
