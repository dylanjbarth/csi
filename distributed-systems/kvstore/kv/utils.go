package kv

import "encoding/binary"

// return byte array prefixed with 4 bytes containing the length of b as uint32
func toWireFormat(b *[]byte) *[]byte {
	l := len(*b)
	out := make([]byte, LENGTH_PREFIX_BYTES)
	binary.LittleEndian.PutUint32(out, uint32(l))
	out = append(out, *b...)
	return &out
}
