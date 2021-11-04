package kv

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
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

func (r *Request) PrettyPrint() string {
	return fmt.Sprintf("Command: %s; Key: %s; Value: %s", r.Command, r.Item.Key, r.Item.Value)
}

func (r *Response) PrettyPrint() string {
	return fmt.Sprintf("Code: %s; Message: %s", r.Code, r.Message)
}

func SendRequest(req *Request, port string) (*Response, error) {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to leader server %s", err)
	}
	reqbytes, err := proto.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize request: %s", err)
	}
	_, err = conn.Write(*toWireFormat(&reqbytes))
	if err != nil {
		return nil, fmt.Errorf("failed to send input to server: %s", err)
	}
	data, err := getNextMessage(&conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from server: %s", err)
	}
	var resp Response
	err = proto.Unmarshal(*data, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize response from server: %s", err)
	}
	return &resp, nil

}

func SendResponse(conn *net.Conn, resp *Response) {
	out, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalf("failed to serialize response: %s", err)
	}
	_, err = (*conn).Write(*toWireFormat(&out))
	if err != nil {
		log.Fatalf("failed to respond to client: %s", err)
	}
}
