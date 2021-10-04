package storage

import (
	"encoding/binary"
	"log"
	"os"
)

type StorageReader struct {
	fp     *os.File
	header *PageHeader
	index  []*IndexEntry
	next   int
}

func NewStorageReader(path string) *StorageReader {
	fp, err := os.Open(path)
	if err != nil {
		log.Fatalf("StorageReader failed to open %s for reading. Error: %s", path, err)
	}
	sw := StorageReader{
		fp: fp,
	}
	sw.readHeader() // read in header and indexes for quick access
	return &sw
}

// Next returns the next record or empty if none remain.
func (sr *StorageReader) Next() string {
	if sr.next >= len(sr.index) {
		return ""
	}
	entry := sr.index[sr.next]
	b := make([]byte, (*entry).DataLen)
	_, err := sr.fp.ReadAt(b, int64((*entry).DataStart))
	if err != nil {
		log.Fatalf("StorageReader failed to read data entry %d. Error: %s", sr.next, err)
	}
	sr.next++
	return string(b)
}

func (sr *StorageReader) readHeader() {
	b := make([]byte, 4) // header
	_, err := sr.fp.ReadAt(b, 0)
	if err != nil {
		log.Fatalf("StorageReader failed to read header. Error: %s", err)
	}
	sr.header = &PageHeader{IndexSize: binary.BigEndian.Uint16(b[:2]), DataSize: binary.BigEndian.Uint16(b[2:])}
	ib := make([]byte, sr.header.IndexSize) // index buffer
	_, err = sr.fp.ReadAt(ib, HeaderLen)
	if err != nil {
		log.Fatalf("StorageReader failed to read indices. Error: %s", err)
	}
	for i := 0; i < int(sr.header.IndexSize); i = i + IndexLen {
		sr.index = append(sr.index, &IndexEntry{DataStart: binary.BigEndian.Uint16(ib[i : i+2]), DataLen: binary.BigEndian.Uint16(ib[i+2:])})
	}
}
