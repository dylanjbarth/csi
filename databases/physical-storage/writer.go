package storage

import (
	"encoding/binary"
	"log"
	"os"
)

type StorageWriter struct {
	fp       *os.File
	colCount int
	file     *File
}

func NewStorageWriter(path string, nCols int) *StorageWriter {
	fp, err := os.Create(path)
	if err != nil {
		log.Fatalf("StorageWriter failed to create %s for writing. Error: %s", path, err)
	}
	// treat file as empty, re-opening a previously written to file out of scope.
	sw := StorageWriter{
		fp:       fp,
		colCount: nCols,
		file: &File{
			Header: PageHeader{
				IndexEntrySize: uint16(2*nCols) + 2, // 2 bytes per column plus data size
			},
		},
	}
	sw.writeHeader() // start with empty
	return &sw
}

func (sw *StorageWriter) Write(t Tuple) bool {
	// TODO: how to serialize a tuple cleanly, how to worry about data type information? Let's assume just strings for simplicity.
	/*
		1. TODO Calculate free space remaining - error if no space
		2. Write tuple index, write tuple, update page header with new offsets.
		3. TODO ideally this would be an atomic write instead of 3 syscalls.
		4. TODO lots of small writes, would be more efficient to use bufio
	*/
	dataSize := 0
	// Calculate size of just the values
	for _, v := range t.Values {
		dataSize += binary.Size([]byte(v.StringValue))
	}

	// Write the data, creating index entries as we go
	indexEntry := IndexEntry{DataSize: uint16(dataSize), ColumnOffsets: []uint16{}}
	start := uint16(FileSize - dataSize - int(sw.file.Header.DataSize))
	offset := 0
	for _, v := range t.Values {
		dPos := start + uint16(offset)
		_, err := sw.fp.WriteAt([]byte(v.StringValue), int64(dPos))
		if err != nil {
			log.Fatalf("StorageWriter failed to write tuple value %s. Error: %s", v.StringValue, err)
		}
		indexEntry.ColumnOffsets = append(indexEntry.ColumnOffsets, dPos)
		offset += binary.Size([]byte(v.StringValue))
	}

	// Update our index
	indexPos := HeaderLen + sw.file.Header.IndexTotalSize
	sw.file.Index = append(sw.file.Index, indexEntry)
	b := make([]byte, sw.file.Header.IndexEntrySize)
	binary.BigEndian.PutUint16(b[0:], uint16(dataSize))
	for i, v := range indexEntry.ColumnOffsets {
		binary.BigEndian.PutUint16(b[i*2+2:i*2+4], v)
	}
	_, err := sw.fp.WriteAt(b, int64(indexPos))
	if err != nil {
		log.Fatalf("StorageWriter failed to update index. Error: %s", err)
	}

	// Update page header:
	sw.file.Header.IndexTotalSize += uint16(len(indexEntry.ColumnOffsets)*2) + 2
	sw.file.Header.DataSize += uint16(dataSize)
	sw.writeHeader()
	return false
}

func (sw *StorageWriter) Close() {
	err := sw.fp.Close()
	if err != nil {
		log.Fatalf("StorageWriter failed to close the file. Error: %s", err)
	}
}

func (sw *StorageWriter) writeHeader() {
	b := make([]byte, HeaderLen)
	binary.BigEndian.PutUint16(b[:2], (*sw.file).Header.IndexEntrySize)
	binary.BigEndian.PutUint16(b[2:4], (*sw.file).Header.IndexTotalSize)
	binary.BigEndian.PutUint16(b[4:], (*sw.file).Header.DataSize)
	_, err := sw.fp.WriteAt(b, 0)
	if err != nil {
		log.Fatalf("StorageWriter failed to write header. Error: %s", err)
	}
}
