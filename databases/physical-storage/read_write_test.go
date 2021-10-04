package storage

import "testing"

func TestWriterAndReader(t *testing.T) {
	// Write a few records to a file
	path := "./test_db"
	sw := NewStorageWriter(path, 2)
	tups := []*Tuple{makeTuple("Hello", "World"), makeTuple("hola", "mundo"), makeTuple("Salut", "le Monde")}
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
	expected = 3 * 6 // data size uint16 * number & size of columns
	if sr.header.IndexTotalSize != uint16(expected) {
		t.Errorf("Expected index offset size of %d but got %d", expected, sr.header.IndexTotalSize)
	}
	expected = 3
	if len(sr.index) != 3 {
		t.Errorf("Expected index len of %d but got %d", expected, len(sr.index))
	}

	first := sr.Next()
	expectedStr := []string{"Hello", "World"}
	for i, s := range expectedStr {
		if s != first[i] {
			t.Errorf("Expected first tuple returned to contain %s at index %d but got %s", s, i, first[i])
		}
	}
	second := sr.Next()
	expectedStr = []string{"hola", "mundo"}
	for i, s := range expectedStr {
		if s != second[i] {
			t.Errorf("Expected second tuple returned to contain %s at index %d but got %s", s, i, second[i])
		}
	}
	third := sr.Next()
	expectedStr = []string{"Salut", "le Monde"}
	for i, s := range expectedStr {
		if s != third[i] {
			t.Errorf("Expected third tuple returned to contain %s at index %d but got %s", s, i, third[i])
		}
	}
}

func makeTuple(uck, name string) *Tuple {
	return &Tuple{
		Values: []Value{{"uck", uck}, {"name", name}},
	}
}
