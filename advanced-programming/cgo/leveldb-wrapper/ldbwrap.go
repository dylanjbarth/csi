package ldbwrap

// #cgo CFLAGS: -I/usr/local/Cellar/leveldb/1.23/include
// #cgo LDFLAGS: -L/usr/local/Cellar/leveldb/1.23/lib -lleveldb
// #include "leveldb/c.h"
import "C"
import "fmt"

func Init() {
	// TODO why are these generated as empty structs by cgo?
	// var db *C.leveldb_t  // cannot use _cgo2 (type **_Ctype_struct_leveldb_t) as type **_Ctype_char in argument to _Cfunc_leveldb_open
	db := C.CString("test")
	opts := C.leveldb_options_t{} // ./ldbwrap.go:13:2: _Ctype_struct_leveldb_options_t is incomplete (or unallocatable); stack allocation disallowed
	// this isn't allowed since this is an empty struct...
	// opts.create_if_missing = true
	status := C.leveldb_open(&opts, C.CString("/tmp/test_level_db"), &db)
	fmt.Println(status)
	// if !status.ok() {
	// 	panic("Unable to open database")
	// }
}

func Open() {

}
