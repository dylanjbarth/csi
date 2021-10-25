package main

import "kv"

const PROMPT = "kv> "

func main() {
	c = kv.NewClient(PROMPT)
	c.DoRepl()
}
