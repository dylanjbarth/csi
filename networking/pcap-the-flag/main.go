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

// https://en.wikipedia.org/wiki/Ethernet_frame
type ethernetHeader struct {
	MacDest [6]byte
	MacSrc  [6]byte
	// The EtherType field is two octets long and it can be used for two different purposes. Values of 1500 and below mean that it is used to indicate the size of the payload in octets, while values of 1536 and above indicate that it is used as an EtherType, to indicate which protocol is encapsulated in the payload of the frame. When used as EtherType, the length of the frame is determined by the location of the interpacket gap and valid frame check sequence (FCS).
	// https://en.wikipedia.org/wiki/EtherType =>
	Ethertype [2]byte
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
	eLen := int(unsafe.Sizeof(ethernetHeader{}))
	readBytesAt(&b, 0, pcapLen, &pcap)
	log.Printf("Magic number is %d, %x", pcap.MagicN, pcap.MagicN)
	log.Printf("Major v: %d, Minor v: %d", pcap.MajorV, pcap.MinorV)
	log.Printf("TzOffset and TsAccuracy should be 0: %d, %d", pcap.TzOffset, pcap.TsAccur)
	log.Printf("Total snapshot byte len %d, %x", pcap.SnapLen, pcap.SnapLen)
	log.Printf("Link Layer Header %d", pcap.LinkHead)

	// Find all the packets
	pOffset := pcapLen
	for i := 1; pOffset < len(b); i++ {
		// Read packet header
		p := packetHeader{}
		readBytesAt(&b, pOffset, pLen+pOffset, &p)
		log.Printf("Packet %d: %ds %dns len: %d, un-truncated len: %d, packet header: %x", i, p.TimestampS, p.TimestampMs, p.Len, p.UntruncLen, b[pOffset:pLen+pOffset])

		// Read ethernet header
		e := ethernetHeader{}
		eOffset := pLen + pOffset // start where the packet header ends
		readBytesAt(&b, eOffset, eLen+eOffset, &e)
		log.Printf("Ethernet frame header: Dest %x Src %x Ethertype %x", e.MacDest, e.MacSrc, e.Ethertype)

		// Before returning, set the new offset for the next packet
		pOffset += int(p.Len) + pLen
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
