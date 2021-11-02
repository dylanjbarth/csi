package kv

import (
	"fmt"
	"strings"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) (Request, error) {
	sp := strings.Fields(input)
	if len(sp) != 2 {
		return Request{}, fmt.Errorf("unable to parse input %s", input)
	}
	cmd := strings.ToLower(sp[0])
	switch cmd {
	case "get":
		return Request{Command: Request_GET, Item: &Item{Key: sp[1]}}, nil
	case "set":
		args := strings.Split(sp[1], "=")
		if len(args) != 2 {
			return Request{}, fmt.Errorf("unable to parse input %s", input)
		}
		return Request{Command: Request_SET, Item: &Item{Key: args[0], Value: args[1]}}, nil
	}
	return Request{}, fmt.Errorf("unable to parse input %s", input)
}
