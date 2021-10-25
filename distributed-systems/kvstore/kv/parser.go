package kv

import (
	"fmt"
	"strings"
)

type Command string

const (
	CMD_GET Command = "get"
	CMD_SET Command = "set"
)

type CommandGroup struct {
	Cmd  Command
	Args []string
}

type Parser struct {
	input string
}

func NewParser(input string) *Parser {
	return &Parser{input}
}

func (p *Parser) Parse() (CommandGroup, error) {
	sp := strings.Fields(p.input)
	if len(sp) != 2 {
		return CommandGroup{}, fmt.Errorf("unable to parse input %s", p.input)
	}
	cmd := strings.ToLower(sp[0])
	switch cmd {
	case "get":
		return CommandGroup{CMD_GET, sp[1:]}, nil
	case "set":
		args := strings.Split(sp[1], "=")
		if len(args) != 2 {
			return CommandGroup{}, fmt.Errorf("unable to parse input %s", p.input)
		}
		return CommandGroup{CMD_SET, args}, nil
	}
	return CommandGroup{}, fmt.Errorf("unable to parse input %s", p.input)
}
