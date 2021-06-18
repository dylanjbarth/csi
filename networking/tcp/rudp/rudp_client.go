package rudp

import (
	"log"
	"syscall"
)

type RUDPClient struct {
	Name       string
	SockPort   int
	SendToPort int

	fd int
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
	return syscall.Sendto(s.fd, b, 0, &syscall.SockaddrInet4{Port: s.SendToPort, Addr: [4]byte{127, 0, 0, 1}})
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
