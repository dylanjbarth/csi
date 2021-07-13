package main

import (
	"testing"
)

func TestSkipListBasicInOrderLevel1(t *testing.T) {
	sl := newSkipListOC(1)
	cases := []string{"a", "b", "c"}
	for _, c := range cases {
		sl.Put(c, c)
	}
}

func TestSkipListBasicInOrderMultiLevels(t *testing.T) {
	sl := newSkipListOC(16)
	cases := []string{"a", "b", "c", "d", "e", "f", "g"}
	for _, c := range cases {
		sl.Put(c, c)
	}
}

func TestSkipList(t *testing.T) {
	// test create, get, update, insert into skip list.
	sl := newSkipListOC(1)
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

	ok = sl.Delete("notthere")
	if ok {
		t.Errorf("Shouldn't have been able to remove non-existent key notthere")
	}
	ok = sl.Delete("z")
	if !ok {
		t.Errorf("Shouldn have been able to remove key z")
	}
	val, ok = sl.Get("z")
	if ok || val != "" {
		t.Errorf("Shouldn't have been able to get key z after delete")
	}
}
