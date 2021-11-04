package kv

const DATA_PATH = "./data/db.kv"
const ROUTER_PORT = ":8010"
const LEADER_PORT = ":8011"
const FOLLOWER_PORT = ":8012"
const LENGTH_PREFIX_BYTES = 4 // for over the wire comms, prefix is uint32, 4 bytes.
