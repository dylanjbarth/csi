package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"unsafe"
)

type pcapHeader struct {
	MagicN   uint32
	MajorV   uint16
	MinorV   uint16
	TzOffset uint32
	TsAccur  uint32
	SnapLen  uint32
	LinkHead uint32
}

type packetHeader struct {
	TimestampS  uint32
	TimestampMs uint32
	Len         uint32
	UntruncLen  uint32
}

func main() {
	b, rerr := ioutil.ReadFile("./net.cap")
	if rerr != nil {
		log.Fatal("Failed to read bytes from network capture.", rerr)
	}
	// file header
	pcap := pcapHeader{}
	pcapLen := int(unsafe.Sizeof(pcap))
	pLen := int(unsafe.Sizeof(packetHeader{}))
	readBytesAt(&b, 0, pcapLen, &pcap)
	log.Printf("Magic number is %d, %x", pcap.MagicN, pcap.MagicN)
	log.Printf("Major v: %d, Minor v: %d", pcap.MajorV, pcap.MinorV)
	log.Printf("TzOffset and TsAccuracy should be 0: %d, %d", pcap.TzOffset, pcap.TsAccur)
	log.Printf("Total snapshot byte len %d, %x", pcap.SnapLen, pcap.SnapLen)
	log.Printf("Link Layer Header %d", pcap.LinkHead)

	// Find all the packets
	offset := pcapLen
	for i := 1; offset < len(b); i++ {
		p := packetHeader{}
		readBytesAt(&b, offset, pLen+offset, &p)
		log.Printf("Packet %d: %ds %dns len: %d, un-truncated len: %d", i, p.TimestampS, p.TimestampMs, p.Len, p.UntruncLen)
		offset += int(p.Len) + pLen
	}
}

func readBytesAt(b *[]byte, start int, end int, d interface{}) {
	bslice := (*b)[start:end]
	// log.Printf("Data as hex %x", bslice)
	breader := bytes.NewReader(bslice)
	err := binary.Read(breader, binary.LittleEndian, d)
	if err != nil {
		log.Fatal("Unable to read bytes: ", err)
	}
}
