package main

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	// test create, get, update, insert into skip list.
	sl := newSkipListOC()
	cases := []string{"z", "a", "b", "c", "y", "d"}
	for _, c := range cases {
		sl.Put(c, c)
	}

	zVal := "foo"
	inserted := sl.Put("z", zVal)
	if inserted {
		t.Errorf("Should have updated key z not inserted.")
	}
	inserted = sl.Put("zz", "bar")
	if !inserted {
		t.Errorf("Should have inserted key zz not updated.")
	}

	val, ok := sl.Get("z")
	if !ok {
		t.Errorf("Should have been able to retrieve value of z")
	}
	if val != zVal {
		t.Errorf("Retrieved unexpected value for z. expected %s but got %s", zVal, val)
	}
}
