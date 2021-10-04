package storage

const HeaderLen = 6
const FileSize = 128 // small for testing

type File struct {
	Header PageHeader
	Index  []IndexEntry
	Data   []byte
}

type PageHeader struct {
	IndexEntrySize uint16
	IndexTotalSize uint16
	DataSize       uint16
}

type IndexEntry struct {
	DataSize      uint16
	ColumnOffsets []uint16
}
