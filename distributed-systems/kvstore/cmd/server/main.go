package main

import (
	"kvstore/kv"
)

func main() {
	// TODO add cli options to determine number of followers.
	leader := kv.NewLeader(kv.DATA_PATH)
	kv.AcceptClientConnections(leader)
}
