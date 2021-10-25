package kv

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Client struct {
	prompt string
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewClient(prompt string) *Client {
	return &Client{prompt, bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)}
}

func (c *Client) DoRepl() {
	for {
		c.writer.WriteString(c.prompt)
		line, err := c.reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Unable to read from stdin. %s", err)
		}
		c.writer.WriteString(fmt.Sprintf("You wrote: %s", line))
	}
}
