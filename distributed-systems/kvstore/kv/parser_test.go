package kv

import "testing"

func TestParser(t *testing.T) {
	p := NewParser()
	out, err := p.Parse("get foo")
	if err != nil {
		t.Errorf("Expected parse to not fail but got %s", err)
	}
	if out.Command != Request_GET {
		t.Errorf("Expected parsed command to be get but got %s", out.Command)
	}
	if out.Item.Key != "foo" {
		t.Errorf("Expected parsed command to be foo but got %s", out.Item.Key)
	}

	out, err = p.Parse("set foo=bar")
	if err != nil {
		t.Errorf("Expected parse to not fail but got %s", err)
	}
	if out.Command != Request_SET {
		t.Errorf("Expected parsed command to be set but got %s", out.Command)
	}
	if out.Item.Key != "foo" {
		t.Errorf("Expected parsed arg to be foo but got %s", out.Item.Key)
	}
	if out.Item.Value != "bar" {
		t.Errorf("Expected parsed arg to be bar but got %s", out.Item.Value)
	}

	out, err = p.Parse("foo")
	if err == nil {
		t.Errorf("Expected parse to fail but got %s", &out)
	}
}
