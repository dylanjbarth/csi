package ldbwrap

// #cgo CFLAGS: -I/usr/local/Cellar/leveldb/1.23/include
// #cgo LDFLAGS: -L/usr/local/Cellar/leveldb/1.23/lib -lleveldb
// #include <stdlib.h>
// #include "leveldb/c.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func Open(path string) {
	errptr := C.CString("")
	opts := C.leveldb_options_create()
	C.leveldb_options_set_create_if_missing(opts, C.uchar(1))
	C.leveldb_open(opts, C.CString(path), &errptr)
	if C.GoString(errptr) != "" {
		panic(C.GoString(errptr))
	} else {
		fmt.Printf("Connected to db %s", path)
	}
	C.free(unsafe.Pointer(errptr))
}
