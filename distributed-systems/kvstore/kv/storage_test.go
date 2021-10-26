package kv

import "testing"

func TestStorage(t *testing.T) {
	s := NewStorage("./test_db.json", true)
	res, err := s.Get("foo")
	if err != nil {
		t.Errorf("Expected get to not fail but got %s", err)
	}
	if res != "" {
		t.Errorf("Expected test db to be empty but got %s", res)
	}

	err = s.Set("foo", "bar")
	if err != nil {
		t.Errorf("Expected set to not fail but got %s", err)
	}

	res, err = s.Get("foo")
	if err != nil {
		t.Errorf("Expected get to not fail but got %s", err)
	}
	if res != "bar" {
		t.Errorf("Expected foo to contain bar but got %s", res)
	}
}
