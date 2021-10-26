package main

import (
	"kvstore/kv"
)

func main() {
	c := kv.NewServer(kv.DATA_PATH)
	c.AcceptConnections()
}
