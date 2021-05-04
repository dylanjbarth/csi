package concurrency

import (
	"reflect"
	"testing"
)

var outer = 1000
var inner = 20
var max = outer * inner

func TestConcurrencyStrategies(t *testing.T) {
	var start uint64
	cases := []counterService{
		// &NoSync{start},  // this fails as expected, TODO is there a nice way to wrap this as an expected failure?
		&AtomicCount{start},
	}
	for _, c := range cases {
		// last := start
		results := make(chan uint64, max)
		// Test monotonically increasing by calling multiple times __within__ a goroutine
		for i := 0; i < outer; i++ {
			go func() {
				var last uint64
				for j := 0; j < inner; j++ {
					new := c.getNext()
					if last > 0 && new < last {
						t.Errorf("Testing %s failed. Values didn't monotonically increase. Received %d after %d within the same goroutine.", reflect.TypeOf(c), new, last)
					}
					last = new
					results <- new
				}
			}()
		}
		// drain channel once tests have run and calc max
		var curr uint64
		for i := 0; i < max; i++ {
			n := <-results
			if n > curr {
				curr = n
			}
		}
		if curr != uint64(max) {
			t.Errorf("Testing %s failed. Expected max value to be %d but got %d.", reflect.TypeOf(c), max, curr)
		}
	}
}
