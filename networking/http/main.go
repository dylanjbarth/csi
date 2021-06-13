package main

import (
	"log"
	"syscall"
)

const maxConn, srcPort, dstPort = 10, 8001, 8002

func main() {
	// create socket to listen on
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("unable to create socket. %s", err)
	}
	// bind socket to source port
	err = syscall.Bind(fd, &syscall.SockaddrInet4{Port: srcPort, Addr: [4]byte{127, 0, 0, 1}})
	if err != nil {
		log.Fatalf("unable to bind socket. %s", err)
	}

	destServer := syscall.SockaddrInet4{Port: dstPort, Addr: [4]byte{127, 0, 0, 1}}

	// set up socket for communication with destination server
	sfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("unable to create socket. %s", err)
	}
	// establish a connection to our destination server
	err = syscall.Connect(sfd, &destServer)
	if err != nil {
		log.Fatalf("unable to connect to dst server. %s", err)
	}

	// Set socket up to listen for connections
	err = syscall.Listen(fd, maxConn)
	if err != nil {
		log.Fatalf("listen failed. %s", err)
	}

	// accept first incoming connection
	nfd, addr, aerr := syscall.Accept(fd)
	if aerr != nil {
		log.Fatalf("accept failed. %s", aerr)
	}
	log.Printf("Accepted conn from %+v", addr)

	for {
		b := make([]byte, 1500) //https://stackoverflow.com/a/2614188 assuming ethernet packet here
		n, frm, rerr := syscall.Recvfrom(nfd, b, 0)
		if rerr != nil {
			log.Fatalf("recvfrom failed. %s", rerr)
		}
		log.Printf("%d, %+v", n, frm)
		log.Printf("%s", b)
		log.Printf("Sending these bytes to our dest server")
		serr := syscall.Sendto(sfd, b, 0, &destServer)
		if serr != nil {
			log.Fatalf("send failed. %s", serr)
		}
	}
}
