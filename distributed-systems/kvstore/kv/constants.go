package kv

const DATA_PATH = "./data/db.kv"
const CLIENT_SERVER_SOCKET_FILE = "kvstore.sock"
const SERVER_SERVER_SOCKET_FILE = "kvstore.serveronly.sock"
const LENGTH_PREFIX_BYTES = 4 // for over the wire comms, prefix is uint32, 4 bytes.
