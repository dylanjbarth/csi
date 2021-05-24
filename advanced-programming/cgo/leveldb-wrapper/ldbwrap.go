package ldbwrap

// #cgo CFLAGS: -I/usr/local/Cellar/leveldb/1.23/include
// #cgo LDFLAGS: -L/usr/local/Cellar/leveldb/1.23/lib -lleveldb
// #include "leveldb/c.h"
import "C"
import "fmt"

func Open() {
	db := C.CString("test")
	opts := C.leveldb_options_create()
	C.leveldb_options_set_create_if_missing(opts, C.uchar(1))
	status := C.leveldb_open(opts, C.CString("/tmp/test_level_db"), &db)
	fmt.Println(*status)
}
