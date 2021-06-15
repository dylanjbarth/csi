package main

import (
	"bytes"
	"fmt"
	"log"
	"syscall"
)

const maxConn, srcPort, dstPort, maxBytes = 10, 8001, 8002, 1500
const cachePath = "cachethis/" // TODO obviously cleaner to allow this to be passed in

func main() {
	emptyBytes := make([]byte, maxBytes)
	// cache := make(map[string][]byte)

	// connect proxy to our end server
	destServer := syscall.SockaddrInet4{Port: dstPort, Addr: [4]byte{127, 0, 0, 1}}
	sfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("unable to create socket. %s", err)
	}
	err = syscall.Connect(sfd, &destServer)
	if err != nil {
		log.Fatalf("unable to connect to dst server. %s", err)
	}

	// create socket to listen for clients on and bind it to our source port
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("unable to create socket. %s", err)
	}
	err = syscall.Bind(fd, &syscall.SockaddrInet4{Port: srcPort, Addr: [4]byte{127, 0, 0, 1}})
	if err != nil {
		log.Fatalf("unable to bind socket. %s", err)
	}
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
		b := make([]byte, maxBytes) //https://stackoverflow.com/a/2614188 assuming ethernet packet here
		_, _, rerr := syscall.Recvfrom(nfd, b, 0)
		if rerr != nil {
			log.Fatalf("recvfrom failed. %s", rerr)
		}
		log.Printf("Message received from client: %s", b)
		// The return value will be 0 when the peer has performed an orderly shutdown.
		if bytes.Equal(emptyBytes, b) {
			log.Print("Got empty response from client, indicating orderly shutdown.")
			break
		}
		val := fmt.Sprintf("%s", b)
		if val != "" {
			log.Printf("Sending these bytes to our dest server")
			serr := syscall.Sendto(sfd, b, 0, &destServer)
			if serr != nil {
				log.Fatalf("send failed. %s", serr)
			}
			// Get the response and forward it back to the original requester
			_, _, rerr = syscall.Recvfrom(sfd, b, 0)
			if rerr != nil {
				log.Fatalf("recvfrom failed. %s", rerr)
			}
			log.Printf("Response from dest server: %s", b)
			log.Printf("Sending these bytes to our client")
			serr = syscall.Sendto(nfd, b, 0, addr)
			if serr != nil {
				log.Fatalf("send failed. %s", serr)
				// TODO getting a broken pipe error here when the client closes the connection..
				// but the bytes are getting through to the client. some flag I should be setting on the sendto?
			}
		} else {
			log.Printf("Message from client was empty so closing the connection.")
			err = syscall.Shutdown(nfd, syscall.SHUT_RDWR)
			if err != nil {
				log.Fatalf("shutdown failed. %s", err)
			}
		}
	}
}
