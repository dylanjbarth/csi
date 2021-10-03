package storage

import "testing"

func TestWriter(t *testing.T) {
	// make sure we can create a storage writer without failing.``
	sw := NewStorageWriter("./test_db")
	tups := []*Tuple{makeTuple("test1", "Hello"), makeTuple("test2", "World")}
	for _, t := range tups {
		sw.Write(*t)
	}
	sw.Close()
}

func makeTuple(uck, name string) *Tuple {
	return &Tuple{
		Values: []Value{{"uck", uck}, {"name", name}},
	}
}
