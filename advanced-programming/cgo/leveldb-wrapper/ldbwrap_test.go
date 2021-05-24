package ldbwrap

import (
	"os"
	"testing"
)

func TestLevelDB(t *testing.T) {

	path := "./test_level_db"
	t.Cleanup(func() {
		os.RemoveAll(path)
	})

	// Open & closing
	ldb := NewLevelDB(path)
	if ldb.isOpen {
		t.Error("Expected new levelDB to be closed after init and before open.")
	}
	ldb.Open()
	if !ldb.isOpen {
		t.Error("Expected levelDB to be open after calling open.")
	}
	ldb.Close()
	if ldb.isOpen {
		t.Error("Expected levelDB to be closed after calling close.")
	}

	// Reopen for more testing
	ldb.Open()

	// Putting and Getting
	k := "foo"
	v := "bar"
	res, getErr := ldb.Get(k)
	if getErr != nil {
		t.Errorf("Expected Get(%s) to not throw but got Error: %s.", k, getErr)
	}
	if res != "" {
		t.Errorf("Expected %s to not have any values stored.", k)
	}
	putErr := ldb.Put(k, v)
	if putErr != nil {
		t.Errorf("Expected Put(%s, %s) to not throw but got Error: %s.", k, v, putErr)
	}

	res, getErr = ldb.Get(k)
	if getErr != nil {
		t.Errorf("Expected Get(%s) to not throw but got Error: %s.", k, getErr)
	}
	if res != v {
		t.Errorf("Expected %s to return %s after we stored it but got %s", k, v, res)
	}
}
