package rudp

import (
	"bytes"
	"encoding/gob"
	"log"
	"syscall"
)

type RUDPClient struct {
	Name       string
	SockPort   int
	SendToPort int

	fd int
}

type RUDPSegment struct {
	Checksum   uint32
	Contentlen uint32
	Body       []byte
}

func (s *RUDPClient) OpenSocket() error {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		return err
	}
	s.fd = fd
	log.Printf("%s has socket %d", s.Name, s.fd)
	return nil
}

func (s *RUDPClient) BindPort() error {
	log.Printf("Binding %s to 127.0.0.1:%d", s.Name, s.SockPort)
	return syscall.Bind(s.fd, &syscall.SockaddrInet4{Port: s.SockPort, Addr: [4]byte{127, 0, 0, 1}})
}

func (s *RUDPClient) Close() error {
	log.Printf("Closing socket for %s", s.Name)
	return syscall.Close(s.fd)
}

func (s *RUDPClient) Send(b []byte) error {
	log.Printf("%s sending bytes '%s' to 127.0.0.1:%d", s.Name, b, s.SendToPort)
	var payload bytes.Buffer
	g := gob.NewEncoder(&payload)
	l := uint32(len(b))
	ck := CalcChecksum(b)
	err := g.Encode(RUDPSegment{ck, l, b})
	if err != nil {
		log.Fatalf("failed to encode rudp packet. %s", err)
	} else {
		log.Printf("Fully encoded segment %b, %+v", payload.Bytes(), payload.Bytes())
	}
	return syscall.Sendto(s.fd, payload.Bytes(), 0, &syscall.SockaddrInet4{Port: s.SendToPort, Addr: [4]byte{127, 0, 0, 1}})
}

func (s *RUDPClient) Receive() ([]byte, int, error) {
	r := make([]byte, 1500)
	log.Printf("%s receiving bytes on socket %d", s.Name, s.fd)
	n, _, err := syscall.Recvfrom(s.fd, r, syscall.MSG_DONTWAIT)
	return r, n, err
}

func (s *RUDPClient) WaitUntilReady() error {
	log.Printf("%s using select syscall to wait until ready on socket %d", s.Name, s.fd)
	fdset := &syscall.FdSet{}
	fdset.Bits[s.fd/32] |= 1 << (uint(s.fd) % 32) // got a little stuck making bitset but found: https://play.golang.org/p/LOd7q3aawd
	return syscall.Select(s.fd+1, fdset, nil, nil, &syscall.Timeval{Sec: 1})
}

func CalcChecksum(bytes []byte) uint32 {
	// https://en.wikipedia.org/wiki/Longitudinal_redundancy_check
	lrc := byte(0)
	for _, b := range bytes {
		lrc ^= b
	}
	return uint32(lrc)
}
