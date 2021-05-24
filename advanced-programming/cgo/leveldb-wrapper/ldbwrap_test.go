package ldbwrap

import "testing"

func TestOpenClose(t *testing.T) {
	ldb := NewLevelDB("/tmp/test_level_db")
	if ldb.isOpen {
		t.Error("Expected new levelDB to be closed after init and before open.")
	}
	ldb.Open()
	if !ldb.isOpen {
		t.Error("Expected levelDB to be open after calling open.")
	}
	t.Logf("Pointer to db connection is %v", ldb.cptr)
	ldb.Close()
	if ldb.isOpen {
		t.Error("Expected levelDB to be closed after calling close.")
	}
	t.Logf("Pointer to db connection is %v", ldb.cptr)
}

// func TestBadPathOpen(t *testing.T) {
// 	Open("/blah/test_level_db")
// }
