package main

import (
	"log"
	"syscall"
)

const maxConn, port = 10, 8001

// echo server
func main() {
	// create socket to listen on
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("unable to create socket. %s", err)
	}

	// bind socket to port 80 for http listening
	err = syscall.Bind(fd, &syscall.SockaddrInet4{Port: port, Addr: [4]byte{127, 0, 0, 1}})
	if err != nil {
		log.Fatalf("unable to bind socket. %s", err)
	}

	// Set socket up to listen for connections
	err = syscall.Listen(fd, maxConn)
	if err != nil {
		log.Fatalf("listen failed. %s", err)
	}

	// accept incoming connections and reply with what we get forever
	nfd, addr, aerr := syscall.Accept(fd)
	if aerr != nil {
		log.Fatalf("accept failed. %s", aerr)
	}
	log.Printf("Accepted conn from %+v. Created new socket %d", addr, nfd)
	for {
		b := make([]byte, 1500) //https://stackoverflow.com/a/2614188 assuming ethernet packet here
		n, frm, rerr := syscall.Recvfrom(nfd, b, 0)
		if rerr != nil {
			log.Fatalf("recvfrom failed. %s", rerr)
		}
		log.Printf("%d, %+v", n, frm)
		log.Printf("%s", b)
		log.Printf("Sending the same bytes back")
		serr := syscall.Sendto(nfd, b, 0, addr)
		if serr != nil {
			log.Fatalf("send failed. %s", serr)
		}
	}
}
