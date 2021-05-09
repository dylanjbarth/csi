package mutex

import "testing"

func TestFootex(t *testing.T) {
	tok := footex{}

	// test that default state is unlocked
	if tok.locked {
		t.Error("Expected default token to be unlocked")
	}

	// test that sync lock works
	tok.Lock()
	if !tok.locked {
		t.Error("Expected Lock to lock the token.")
	}

	// test that sync unlock works
	tok.Unlock()
	if tok.locked {
		t.Error("Expected Unlock to unlock the token.")
	}

	// Test concurrent access to shared var.
	n_test := 1000
	d := uint32(0)
	results := make(chan uint32, n_test)
	for i := 0; i < n_test; i++ {
		go func() {
			tok.Lock()
			d += 1
			results <- d
			tok.Unlock()
		}()
	}
	// Ensure that we incremented correctly and have the right max
	var max uint32
	for i := 0; i < n_test; i++ {
		r := <-results
		if r > max {
			max = r
		}
	}
	expected := uint32(n_test) - 1
	if max != expected {
		t.Errorf("Expected max value to equal %d but got %d", expected, max)
	}
}
