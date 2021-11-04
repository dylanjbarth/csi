package main

import (
	"kvstore/kv"
)

func main() {
	// TODO add cli options to determine number of followers.
	router := kv.NewRouter()
	go router.AcceptClientConnections()
	leader := kv.NewLeader(kv.DATA_PATH)
	kv.AcceptRouterConnections(leader)
}
