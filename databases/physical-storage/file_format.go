package storage

const HeaderLen = 4
const IndexLen = 4
const FileSize = 256

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
