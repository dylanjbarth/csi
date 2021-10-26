package kv

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Client struct {
	prompt string
	reader *bufio.Reader
	writer io.Writer
}

func NewClient(prompt string) *Client {
	return &Client{prompt, bufio.NewReader(os.Stdin), os.Stdout}
}

func (c *Client) DoRepl() {
	for {
		_, err := c.writer.Write([]byte(c.prompt))
		if err != nil {
			log.Fatalf("failed to write to stdout: %s", err)
		}
		line, err := c.reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Unable to read from stdin. %s", err)
		}
		c.writer.Write([]byte(fmt.Sprintf("You wrote: %s", line)))
	}
}
