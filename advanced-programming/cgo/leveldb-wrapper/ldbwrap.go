package ldbwrap

// #cgo LDFLAGS: -L/usr/local/Cellar/leveldb/1.23/lib
// #include "leveldb/c.h"
import "C"

func Init() {
	var db C.leveldb_t
	var opts C.leveldb_options_t
	opts.create_if_missing = true
	// status := C.leveldb_open(opts, "/tmp/test_level_db", &db)
	// if !status.ok() {
	// 	panic("Unable to open database")
	// }
}

func Open() {

}
