package kv

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/protobuf/proto"
)

type Client struct {
	prompt string
	conn   net.Conn
	reader *bufio.Reader
	writer io.Writer
	p      *Parser
}

func NewClient(prompt string) *Client {
	conn := initConnection()
	return &Client{prompt, conn, bufio.NewReader(os.Stdin), os.Stdout, NewParser()}
}

func initConnection() net.Conn {
	conn, err := net.Dial("unix", SOCKET_FILE)
	if err != nil {
		log.Fatalf("Unable to connect to server. Is it running? %s", err)
	}
	return conn
}

func (c *Client) Shell() {
	for {
		_, err := c.writer.Write([]byte(c.prompt))
		if err != nil {
			log.Fatalf("failed to write to stdout: %s", err)
		}
		line, err := c.reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Unable to read from stdin. %s", err)
		}
		req, err := c.p.Parse(line)
		if err != nil {
			c.RespondToUser(&Response{Code: Response_FAILURE, Message: fmt.Sprintf("failed to parse input %s", err)})
		} else {
			resp := c.SendToServer(&req)
			c.RespondToUser(resp)
		}
	}
}

func (c *Client) SendToServer(req *Request) *Response {
	reqbytes, err := proto.Marshal(req)
	if err != nil {
		log.Fatalf("failed to serialize request: %s", err)
	}
	_, err = c.conn.Write(*toWireFormat(&reqbytes))
	if err != nil {
		log.Fatalf("failed to send input to server: %s", err)
	}
	data, err := getNextMessage(&c.conn)
	if err != nil {
		log.Fatalf("failed to read response from server: %s", err)
	}
	var resp Response
	err = proto.Unmarshal(*data, &resp)
	if err != nil {
		log.Fatalf("failed to deserialize response from server: %s", err)
	}
	return &resp
}

func (c *Client) RespondToUser(res *Response) {
	c.writer.Write([]byte(res.PrettyPrint()))
}
