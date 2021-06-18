package main

import (
	"syscall"
)

type RUDPClient struct {
	fd      int
	port    int
	dstAdd  [4]byte
	dstPort int
}

func (s *RUDPClient) OpenSocket() error {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return err
	}
	s.fd = fd
	return nil
}

func (s *RUDPClient) BindPort() error {
	return syscall.Bind(s.fd, &syscall.SockaddrInet4{Port: s.port, Addr: [4]byte{127, 0, 0, 1}})
}

func (s *RUDPClient) Close() error {
	return syscall.Close(s.fd)
}

func (s *RUDPClient) Send(b []byte) error {
	return syscall.Sendto(s.fd, b, 0, &syscall.SockaddrInet4{Port: s.dstPort, Addr: s.dstAdd})
}

func (s *RUDPClient) Receive() ([]byte, int, error) {
	r := make([]byte, 1500)
	n, _, err := syscall.Recvfrom(s.fd, r, 0)
	return r, n, err
}
