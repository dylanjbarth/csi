package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
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

// https://datatracker.ietf.org/doc/html/rfc791
type ipHeader struct {
	VersionAndIHL          byte
	TypeOfService          byte
	TotalLength            uint16
	Identification         [2]byte
	FlagsAndFragmentOffset [2]byte
	TTL                    byte
	Protocol               byte
	HeaderChecksum         [2]byte
	SrcAddress             [4]byte
	DestAddress            [4]byte
}

// https://en.wikipedia.org/wiki/Transmission_Control_Protocol#TCP_segment_structure
type tcpHeader struct {
	SrcPort      uint16
	DestPort     uint16
	SeqNum       uint32
	AckNum       uint32
	DataAndFlags [2]byte
	WindowSize   uint16
	Checksum     uint16
	UrgentPtr    uint16
}

const dataDir = "./data/"

func main() {
	// clean up data dir if it exists and re-create it
	err := os.RemoveAll(dataDir)
	if err != nil {
		log.Fatalf("Unable to remove files from data dir. %s", err)
	}
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err := os.Mkdir(dataDir, os.FileMode(0700))
		if err != nil {
			log.Fatalf("Unable to create data dir for HTTP segments. %s", err)
		}
	}

	b, rerr := ioutil.ReadFile("./net.cap")
	if rerr != nil {
		log.Fatal("Failed to read bytes from network capture.", rerr)
	}

	// file header
	pcap := pcapHeader{}
	pcapLen := int(unsafe.Sizeof(pcap))
	pLen := int(unsafe.Sizeof(packetHeader{}))
	eLen := int(unsafe.Sizeof(ethernetHeader{}))
	iLen := int(unsafe.Sizeof(ipHeader{}))
	tLen := int(unsafe.Sizeof(tcpHeader{}))
	readBytesAt(&b, 0, pcapLen, &pcap, false)
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
		readBytesAt(&b, pOffset, pLen+pOffset, &p, false)
		log.Printf("Packet %d: %ds %dns len: %d, un-truncated len: %d, packet header: %x", i, p.TimestampS, p.TimestampMs, p.Len, p.UntruncLen, b[pOffset:pLen+pOffset])

		// Read ethernet header
		e := ethernetHeader{}
		eOffset := pLen + pOffset // start where the packet header ends
		readBytesAt(&b, eOffset, eLen+eOffset, &e, false)
		log.Printf("Ethernet frame header: Dest %x Src %x Ethertype %x", e.MacDest, e.MacSrc, e.Ethertype)

		// Read IP header
		ip := ipHeader{}
		iOffset := eOffset + eLen // start where the ethernet header ends
		readBytesAt(&b, iOffset, iLen+iOffset, &ip, false)
		ihl := (ip.VersionAndIHL & 0x0f) * 4
		log.Printf("%x", ip)
		log.Printf("IP frame header: %+v", ip)

		// Read TCP header
		tcp := tcpHeader{}
		tOffset := iOffset + int(ihl) // start where the ip header ends
		// NB: flipping to parse these bytes as big-endian because network packets always use it
		readBytesAt(&b, tOffset, tLen+tOffset, &tcp, true)
		log.Printf("TCP header: %+v", tcp)
		log.Printf("Source port: %0b, %x, %d", tcp.SrcPort, tcp.SrcPort, tcp.SrcPort)
		tcpHeaderLen := 5 * 4 // min data offset
		dataOffset := tcp.DataAndFlags[0] >> 4
		if dataOffset > 5 {
			tcpHeaderLen = int(dataOffset) * 4
		}
		log.Printf("tcpHeaderLen %d", tcpHeaderLen)
		// Total tcp header len = sizeof tcp struct +
		// Total IP datagram len - (TCP header len + IP Header len) = data section len.
		dataLen := int(ip.TotalLength) - (iLen + tcpHeaderLen)

		// dump this data segment
		data := b[tLen+tcpHeaderLen : tLen+tcpHeaderLen+dataLen]
		// log.Printf("Data segment: %s", data)
		// Only store HTTP response data not request data so looking for dest port 80 and flag are set to PSH or ACK
		log.Printf("%0b", tcp.DataAndFlags)
		// pshSet := (tcp.DataAndFlags[1] >> 3) & 1
		// ackSet := (tcp.DataAndFlags[1] >> 4) & 1
		// if tcp.SrcPort == 80 && (pshSet == 1 || ackSet == 1) {
		if tcp.SrcPort == 80 {
			ferr := ioutil.WriteFile(fmt.Sprintf("%s%d.cap", dataDir, tcp.SeqNum), data, fs.FileMode(0600))
			if ferr != nil {
				log.Fatalf("Failed to write data segment. %s", ferr)
			} else {
				log.Printf("Stored http data for packet %d in %d.cap", i, tcp.SeqNum)
			}
		}

		// Before returning, set the new offset for the next packet
		pOffset += int(p.Len) + pLen
	}

	// process the HTTP segments into a single file
	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatalf("Couldn't read files in data dir %s", err)
	}
	var allBytes []byte
	for _, file := range files {
		// these are sorted already thanks to ReadDir
		b, err := ioutil.ReadFile(fmt.Sprintf("%s%s", dataDir, file.Name()))
		if err != nil {
			log.Fatalf("Couldn't read %s. %s", file.Name(), err)
		}
		allBytes = append(allBytes, b...)
	}
	log.Printf("%s\n", allBytes)
	ioutil.WriteFile(fmt.Sprintf("%sout.jpg", dataDir), allBytes, fs.FileMode(0600))
}

func readBytesAt(b *[]byte, start int, end int, d interface{}, bigEndian bool) {
	bslice := (*b)[start:end]
	log.Printf("Bytes[%d:%d] - as hex %x", start, end, bslice)
	breader := bytes.NewReader(bslice)
	var endian binary.ByteOrder = binary.LittleEndian
	if bigEndian {
		endian = binary.BigEndian
	}
	err := binary.Read(breader, endian, d)
	if err != nil {
		log.Fatal("Unable to read bytes: ", err)
	}
}
