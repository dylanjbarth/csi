package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"syscall"
)

const maxConn, srcPort, dstPort, maxBytes = 10, 8001, 8002, 1500
const cachePath = "cachethis/" // TODO obviously cleaner to allow this to be passed in

func main() {
	emptyBytes := make([]byte, maxBytes)
	cache := make(map[string][]byte)

	// create socket to listen for clients on and bind it to our source port
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	exitIfErr(err, "unable to create socket.")
	err = syscall.Bind(fd, &syscall.SockaddrInet4{Port: srcPort, Addr: [4]byte{127, 0, 0, 1}})
	exitIfErr(err, "unable to bind socket.")
	err = syscall.Listen(fd, maxConn)
	exitIfErr(err, "listen failed.")

	// accept connections
	nfd, addr, err := syscall.Accept(fd)
	exitIfErr(err, "accept failed.")
	log.Printf("Accepted conn from %+v", addr)

	for {
		b := make([]byte, maxBytes) //https://stackoverflow.com/a/2614188 assuming ethernet packet size here, 1500 bytes
		_, _, err := syscall.Recvfrom(nfd, b, 0)
		exitIfErr(err, "recvfrom failed.")
		log.Printf("Message received from client: %s", b)

		// The return value will be 0 when the peer has performed an orderly shutdown.
		if bytes.Equal(emptyBytes, b) {
			log.Print("Got empty response from client, indicating orderly shutdown.")
			break
		}

		// determine if we need to cache it
		cacheIt := false
		respLines := strings.Split(fmt.Sprintf("%s", b), "\r\n")
		reqLine := strings.Split(respLines[0], " ")
		reqPath := reqLine[1]
		if strings.Contains(reqPath, cachePath) {
			if val, ok := cache[reqPath]; ok {
				log.Printf("Resp already cached, returning bytes directly to client.")
				err := syscall.Sendto(nfd, val, 0, addr)
				exitIfErr(err, "send failed.")
				break
			}
			cacheIt = true
		}

		// connect proxy to our end server
		dstServer := syscall.SockaddrInet4{Port: dstPort, Addr: [4]byte{127, 0, 0, 1}}
		sfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
		exitIfErr(err, "unable to create socket.")

		err = syscall.Connect(sfd, &dstServer)
		exitIfErr(err, "unable to connect to dst server.")

		log.Printf("Sending these bytes to our dst server")
		err = syscall.Sendto(sfd, b, 0, &dstServer)
		exitIfErr(err, "send failed.")

		// Get the response
		_, _, err = syscall.Recvfrom(sfd, b, 0)
		exitIfErr(err, "recvfrom failed.")
		log.Printf("Response from dst server: %s", b)

		// close connection with dest server
		err = syscall.Close(sfd)
		exitIfErr(err, "close failed.")

		if cacheIt {
			log.Printf("Storing this response in cache: %s", reqPath)
			cache[reqPath] = b
		}

		// finally forward resp back to the original requester
		log.Printf("Sending these bytes to our client")
		err = syscall.Sendto(nfd, b, 0, addr)
		exitIfErr(err, "send failed.")
	}
}

func exitIfErr(e error, msg string) {
	if e != nil {
		log.Fatalf("%s. %s", msg, e)
	}
}
