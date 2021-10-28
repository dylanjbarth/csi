package kv

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

type Client struct {
	prompt string
	conn   net.Conn
	reader *bufio.Reader
	writer io.Writer
}

func NewClient(prompt string) *Client {
	conn := initConnection()
	return &Client{prompt, conn, bufio.NewReader(os.Stdin), os.Stdout}
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
		resp := c.SendToServer(line)
		c.writer.Write([]byte(resp))
	}
}

func (c *Client) SendToServer(line string) string {
	_, err := c.conn.Write([]byte(line))
	if err != nil {
		log.Fatalf("failed to send input to server: %s", err)
	}
	resp, err := bufio.NewReader(c.conn).ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read response from server: %s", err)
	}
	return resp
}
