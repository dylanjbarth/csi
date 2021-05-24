package ldbwrap

// #cgo CFLAGS: -I/usr/local/Cellar/leveldb/1.23/include
// #cgo LDFLAGS: -L/usr/local/Cellar/leveldb/1.23/lib -lleveldb
// #include <stdlib.h>
// #include "leveldb/c.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type LevelDB struct {
	path   string
	isOpen bool
	cptr   *C.leveldb_t
}

func NewLevelDB(path string) *LevelDB {
	return &LevelDB{path: path, isOpen: false}
}

func (ldb *LevelDB) Open() (*LevelDB, error) {
	if ldb.isOpen {
		return nil, errors.New("database is already open")
	}
	errptr := C.CString("")
	opts := C.leveldb_options_create()
	C.leveldb_options_set_create_if_missing(opts, C.uchar(1))
	ldb.cptr = C.leveldb_open(opts, C.CString(ldb.path), &errptr)
	err := C.GoString(errptr)
	C.free(unsafe.Pointer(errptr))
	if err != "" {
		return nil, fmt.Errorf("failed to connect to DB. error is: %s", err)
	}
	ldb.isOpen = true
	return ldb, nil
}

func (ldb *LevelDB) Close() (*LevelDB, error) {
	if !ldb.isOpen {
		return nil, errors.New("database is already closed")
	}
	C.leveldb_close(ldb.cptr)
	ldb.isOpen = false
	ldb.cptr = nil
	return ldb, nil
}

func (ldb *LevelDB) Get(key string) (string, error) {
	if !ldb.isOpen {
		return "", errors.New("database must be opened before operations can be performed")
	}
	errptr := C.CString("")
	opts := C.leveldb_readoptions_create()
	vallen := C.ulong(0)
	res := C.leveldb_get(ldb.cptr, opts, C.CString(key), C.ulong(len(key)), &vallen, &errptr)
	err := C.GoString(errptr)
	C.free(unsafe.Pointer(errptr))
	if err != "" {
		return "", fmt.Errorf("failed to get %s from DB. error is: %s", key, err)
	}
	// NB taking slice here to the length of vallen to trim whitespace.
	return C.GoString(res)[0:vallen], nil
}

func (ldb *LevelDB) Put(key string, val string) error {
	if !ldb.isOpen {
		return errors.New("database must be opened before operations can be performed")
	}
	errptr := C.CString("")
	opts := C.leveldb_writeoptions_create()
	C.leveldb_put(ldb.cptr, opts, C.CString(key), C.ulong(len(key)), C.CString(val), C.ulong(len(val)), &errptr)
	err := C.GoString(errptr)
	C.free(unsafe.Pointer(errptr))
	if err != "" {
		return fmt.Errorf("failed to put %s,%s to DB. error is: %s", key, val, err)
	}
	return nil
}
