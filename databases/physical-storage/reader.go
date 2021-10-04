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
func (sr *StorageReader) Next() []string {
	if sr.next >= len(sr.index) {
		return []string{}
	}
	entry := sr.index[sr.next]
	output := []string{}
	bytesRemaining := (*entry).DataSize
	for i, offset := range (*entry).ColumnOffsets {
		var b []byte
		if i == len((*entry).ColumnOffsets)-1 {
			b = make([]byte, bytesRemaining)
		} else {
			bufSize := (*entry).ColumnOffsets[i+1] - offset
			b = make([]byte, bufSize)
			bytesRemaining = bytesRemaining - bufSize
		}
		_, err := sr.fp.ReadAt(b, int64((*entry).ColumnOffsets[i]))
		if err != nil {
			log.Fatalf("StorageReader failed to read data entry %d. Error: %s", sr.next, err)
		}
		output = append(output, string(b))
	}
	sr.next++
	return output
}

func (sr *StorageReader) readHeader() {
	b := make([]byte, HeaderLen) // header
	_, err := sr.fp.ReadAt(b, 0)
	if err != nil {
		log.Fatalf("StorageReader failed to read header. Error: %s", err)
	}
	sr.header = &PageHeader{IndexEntrySize: binary.BigEndian.Uint16(b[:2]), IndexTotalSize: binary.BigEndian.Uint16(b[2:4]), DataSize: binary.BigEndian.Uint16(b[4:])}

	ib := make([]byte, sr.header.IndexTotalSize) // index buffer
	_, err = sr.fp.ReadAt(ib, HeaderLen)
	if err != nil {
		log.Fatalf("StorageReader failed to read indices. Error: %s", err)
	}

	// iterate by index entry sizes, starting with the first index.
	for index := 0; index < int(sr.header.IndexTotalSize); index = index + int(sr.header.IndexEntrySize) {
		// The first two bytes are the data size, and then the column offsets follow in 2 byte increments.
		entry := IndexEntry{DataSize: binary.BigEndian.Uint16(ib[index : index+2]), ColumnOffsets: []uint16{}}
		for offset := 2; offset < int(sr.header.IndexEntrySize); offset = offset + 2 {
			entry.ColumnOffsets = append(entry.ColumnOffsets, binary.BigEndian.Uint16(ib[index+offset:index+offset+2]))
		}
		sr.index = append(sr.index, &entry)
	}
}
