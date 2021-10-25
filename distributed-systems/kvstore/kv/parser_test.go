package kv

import "testing"

func TestParser(t *testing.T) {
	p := NewParser("get foo")
	out, err := p.Parse()
	if err != nil {
		t.Errorf("Expected parse to not fail but got %s", err)
	}
	if out.Cmd != "get" {
		t.Errorf("Expected parsed command to be get but got %s", out.Cmd)
	}
	if len(out.Args) != 1 {
		t.Errorf("Expected arguments to be len 1 but got %d", len(out.Args))
	}
	if out.Args[0] != "foo" {
		t.Errorf("Expected parsed command to be foo but got %s", out.Args[0])
	}

	p = NewParser("set foo=bar")
	out, err = p.Parse()
	if err != nil {
		t.Errorf("Expected parse to not fail but got %s", err)
	}
	if out.Cmd != "set" {
		t.Errorf("Expected parsed command to be set but got %s", out.Cmd)
	}
	if len(out.Args) != 2 {
		t.Errorf("Expected arguments to be len 2 but got %d", len(out.Args))
	}
	if out.Args[0] != "foo" {
		t.Errorf("Expected parsed arg to be foo but got %s", out.Args[0])
	}
	if out.Args[1] != "bar" {
		t.Errorf("Expected parsed arg to be bar but got %s", out.Args[1])
	}

	p = NewParser("foo")
	out, err = p.Parse()
	if err == nil {
		t.Errorf("Expected parse to fail but got %s", out)
	}
}
