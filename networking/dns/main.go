package main

import (
	"encoding/binary"
	"log"
	"math/rand"
	"strings"
	"syscall"
)

type DNSHeader struct {
	ID      uint16
	Flags   uint16
	QdCount uint16
	AnCount uint16
	NsCount uint16
	ArCount uint16
}

type DNSQuestionName struct {
	Len  uint8
	Name []byte
}

type DNSResourceRecord struct {
	Name     string
	Type     uint16
	Class    uint16
	TTL      uint32
	Rdlength uint16
	Rdata    string
}

type DNSQuery struct {
	Header  DNSHeader
	Queries []DNSResourceRecord
}

func main() {
	// make socket
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		log.Fatalf("failed to open socket. %s", err)
	}
	log.Printf("Got file descriptor %d", fd)

	// bind socket to any port
	berr := syscall.Bind(fd, &syscall.SockaddrInet4{Port: 0, Addr: [4]byte{0, 0, 0, 0}})
	if berr != nil {
		log.Fatalf("failed to bind socket to port. %s", berr)
	}

	// try to send a DNS query to google
	q := DNSQuery{
		Header: DNSHeader{
			ID:      uint16(rand.Uint32()),
			Flags:   0x0120,
			QdCount: 1,
			AnCount: 0,
			NsCount: 0,
			ArCount: 0,
		},
		Queries: []DNSResourceRecord{
			DNSResourceRecord{
				Name:  "dylanbarth.com",
				Type:  1, // A
				Class: 1, // IN
			},
		},
	}

	// serialize from struct to bytes
	b := make([]byte, 12)
	/*
			The header contains the following fields:
		                                    1  1  1  1  1  1
		      0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		    |                      ID                       |
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		    |QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		    |                    QDCOUNT                    |
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		    |                    ANCOUNT                    |
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		    |                    NSCOUNT                    |
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		    |                    ARCOUNT                    |
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/
	binary.BigEndian.PutUint16(b[0:2], q.Header.ID)
	binary.BigEndian.PutUint16(b[2:4], q.Header.Flags)
	binary.BigEndian.PutUint16(b[4:6], q.Header.QdCount)
	binary.BigEndian.PutUint16(b[6:8], q.Header.AnCount)
	binary.BigEndian.PutUint16(b[8:10], q.Header.NsCount)
	binary.BigEndian.PutUint16(b[10:12], q.Header.ArCount)

	/*

			The question section is used to carry the "question" in most queries,
		i.e., the parameters that define what is being asked.  The section
		contains QDCOUNT (usually 1) entries, each of the following format:

		                                    1  1  1  1  1  1
		      0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		    |                                               |
		    /                     QNAME                     /
		    /                                               /
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		    |                     QTYPE                     |
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
		    |                     QCLASS                    |
		    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/
	// split into segments
	parts := strings.Split(q.Queries[0].Name, ".")
	for _, v := range parts {
		// encode the len then each char
		l := len(v)
		bi := append([]byte{byte(l)}, []byte(v)...)
		// add this to the whole request
		b = append(b, bi...)
	}
	// add a null byte for root to finish the Name
	b = append(b, byte(0))
	// b = append(b, 0x00)
	// Create extra space for type and class at end of byte array
	b = append(b, []byte{0, 0, 0, 0}...)
	// add qtype and class
	// l := len(b)
	binary.BigEndian.PutUint16(b[28:30], q.Queries[0].Type)
	binary.BigEndian.PutUint16(b[30:32], q.Queries[0].Class)

	// send the bytes to google
	serr := syscall.Sendto(fd, b, 0, &syscall.SockaddrInet4{Port: 53, Addr: [4]byte{8, 8, 8, 8}})
	if serr != nil {
		log.Fatalf("failed to send query to google. %s", serr)
	}

	// wait for the response
	r := make([]byte, 4096)
	n, from, rerr := syscall.Recvfrom(fd, r, 0)
	if rerr != nil {
		log.Fatalf("failed to receive from google. %s", rerr)
	}
	log.Printf("n: %d, %+v, %v", n, from, r)
}
