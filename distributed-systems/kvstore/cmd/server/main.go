package main

import (
	"kvstore/kv"
)

func main() {
	// TODO add cli options to determine number of followers.
	router := kv.NewRouter()
	go router.AcceptClientConnections()
	leader := kv.NewLeader(kv.LEADER_DATA_PATH)
	go kv.AcceptRouterConnections(leader)
	follower := kv.NewFollower(kv.FOLLOWER_DATA_PATH)
	kv.AcceptRouterConnections(follower)
}
