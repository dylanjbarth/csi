package storage

import "testing"

func TestWriterAndReader(t *testing.T) {
	// Write a few records to a file
	path := "./test_db"
	sw := NewStorageWriter(path)
	tups := []*Tuple{makeTuple("Hello", "World"), makeTuple("bienvenidos", "mundo"), makeTuple("foo", "bar")}
	for _, t := range tups {
		sw.Write(*t)
	}
	sw.Close()

	// Then try and read them back
	sr := NewStorageReader(path)
	expected := 32
	if sr.header.DataSize != uint16(expected) {
		t.Errorf("Expected data size of %d but got %d", expected, sr.header.DataSize)
	}
	expected = IndexLen * 3
	if sr.header.IndexSize != uint16(expected) {
		t.Errorf("Expected index offset size of %d but got %d", expected, sr.header.IndexSize)
	}
	expected = 3
	if len(sr.index) != 3 {
		t.Errorf("Expected index len of %d but got %d", expected, len(sr.index))
	}

	first := sr.Next()
	expectedStr := "HelloWorld"
	if first != expectedStr {
		t.Errorf("Expected index len of %s but got %s", expectedStr, first)
	}
	second := sr.Next()
	expectedStr = "bienvenidosmundo"
	if second != expectedStr {
		t.Errorf("Expected index len of %s but got %s", expectedStr, second)
	}
}

func makeTuple(uck, name string) *Tuple {
	return &Tuple{
		Values: []Value{{"uck", uck}, {"name", name}},
	}
}
