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
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
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
	n, _, err := syscall.Recvfrom(s.fd, r, 0)
	return r, n, err
}
