package kv

const DATA_PATH = "./data/db.kv"
const SOCKET_FILE = "kvstore.sock"
const LENGTH_PREFIX_BYTES = 4 // for over the wire comms, prefix is uint32, 4 bytes.
