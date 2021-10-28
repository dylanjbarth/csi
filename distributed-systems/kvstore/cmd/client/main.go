package main

import (
	"kvstore/kv"
)

const PROMPT = "kv> "

func main() {
	c := kv.NewClient(PROMPT)
	c.Shell()
}
