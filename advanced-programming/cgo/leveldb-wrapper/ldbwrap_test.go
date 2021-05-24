package ldbwrap

import "testing"

func TestOpen(t *testing.T) {
	Open("/tmp/test_level_db")
}

func TestBadPathOpen(t *testing.T) {
	Open("/blah/test_level_db")
}
