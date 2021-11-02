package kv

import (
	"encoding/binary"
	"testing"
)

// return byte array prefixed with 4 bytes containing the length of b as uint32
func TestToWireFormat(t *testing.T) {
	test := []byte{1, 1, 1, 1}
	out := toWireFormat(&test)
	if len(*out) != 8 {
		t.Errorf("expected toWireFormat to prepend with uint32 but got %d bytes", len(*out))
	}
	lenprefix := binary.LittleEndian.Uint32((*out)[:4])
	if lenprefix != LENGTH_PREFIX_BYTES {
		t.Errorf("expected toWireFormat to prepend with length but got %d bytes", lenprefix)
	}

	test = []byte{1, 1, 1, 1, 1, 1, 1}
	out = toWireFormat(&test)
	if len(*out) != 11 {
		t.Errorf("expected toWireFormat to prepend with uint32 but got %d bytes", len(*out))
	}
	lenprefix = binary.LittleEndian.Uint32((*out)[:4])
	if lenprefix != 7 {
		t.Errorf("expected toWireFormat to prepend with length but got %d bytes", lenprefix)
	}
}
