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

}
