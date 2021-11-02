package main

import (
	"kvstore/kv"
)

func main() {
	leader := kv.NewLeader(kv.DATA_PATH)
	kv.AcceptConnections(leader)
}
