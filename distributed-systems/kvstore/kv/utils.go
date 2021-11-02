package kv

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// return byte array prefixed with 4 bytes containing the length of b as uint32
func toWireFormat(b *[]byte) *[]byte {
	l := len(*b)
	out := make([]byte, LENGTH_PREFIX_BYTES)
	binary.LittleEndian.PutUint32(out, uint32(l))
	out = append(out, *b...)
	return &out
}

// return single protobuf message from wire by reading length prefix and then rest of message.
// response is returned as byte array so caller can interpret as they see fit.
func getNextMessage(c *net.Conn) (*[]byte, error) {
	reader := bufio.NewReader(*c)
	len, err := reader.Peek(LENGTH_PREFIX_BYTES)
	if err != nil {
		return nil, fmt.Errorf("failed to read length prefix from message: %s", err)
	}
	l := binary.LittleEndian.Uint32(len)
	out := make([]byte, l+LENGTH_PREFIX_BYTES)
	_, err = io.ReadFull(reader, out)
	if err != nil {
		return nil, fmt.Errorf("failed to read message: %s", err)
	}
	out = out[LENGTH_PREFIX_BYTES:]
	return &out, nil
}
