package ldbwrap

// #cgo CFLAGS: -I/usr/local/Cellar/leveldb/1.23/include
// #cgo LDFLAGS: -L/usr/local/Cellar/leveldb/1.23/lib
// #include "leveldb/c.h"
import "C"

func Init() {
	// TODO why are these empty structs?
	var db C.leveldb_t
	var opts C.leveldb_options_t
	// this isn't allowed since this is an empty struct...
	opts.create_if_missing = true
	/*
		$ go build
		# ldbwrap
		Undefined symbols for architecture x86_64:
			"_leveldb_open", referenced from:
					__cgo_67271468db4f_Cfunc_leveldb_open in _x002.o
				(maybe you meant: __cgo_67271468db4f_Cfunc_leveldb_open)
		ld: symbol(s) not found for architecture x86_64
		clang: error: linker command failed with exit code 1 (use -v to see invocation)
	*/
	status := C.leveldb_open(opts, "/tmp/test_level_db", &db)
	if !status.ok() {
		panic("Unable to open database")
	}
}

func Open() {

}
