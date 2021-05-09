package mutex

import (
	"testing"
)

func TestFootex(t *testing.T) {
	tok := footex{}

	// Test concurrent access to shared var.
	n_test := 1000
	d := uint32(0)
	results := make(chan uint32, n_test)
	for i := 0; i < n_test; i++ {
		// t.Logf("Starting %d\n", i)
		go func(n int) {
			// t.Logf("%d calling lock\n", n)
			tok.Lock()
			d += 1
			results <- d
			// t.Logf("%d calling unlock\n", n)
			tok.Unlock()
			// t.Logf("%d returning\n", n)
		}(i)
	}
	// Ensure that we incremented correctly and have the right max
	var max uint32
	for i := 0; i < n_test; i++ {
		r := <-results
		if r > max {
			max = r
		}
	}
	expected := uint32(n_test)
	if max != expected {
		t.Errorf("Expected max value to equal %d but got %d", expected, max)
	}
}
