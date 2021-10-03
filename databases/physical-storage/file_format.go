package storage

type File struct {
	Header PageHeader
	Index  []IndexEntry
	Data   []byte
}

type PageHeader struct {
	IndexSize uint16
	DataSize  uint16
}

type IndexEntry struct {
	DataStart uint16
	DataLen   uint16
}
