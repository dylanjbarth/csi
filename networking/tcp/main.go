package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"sync"
	"syscall"

	"./rudp"
)

// Sender
// echo 'hello' | nc -u 127.0.0.1 64132
// Receiver
// while true; do nc -l -u 0.0.0.0 9999; done

/*
Example

$ python3 unreliable_proxy.py 8000
Forwarding 0.0.0.0:51948 -> 127.0.0.1:8000

We'd want to start our program like this:

go run main.go --sendPort 51948 --recPort 8000

so that we send bytes to the proxy on 51948 and receive data from the proxy on port 8000
*/

const defaultSendPort, defaultRecPort = 52058, 8000

func main() {
	sendPortPtr := flag.Int("sendPort", defaultSendPort, fmt.Sprintf("port to send on. default %d", defaultSendPort))
	recPortPtr := flag.Int("recPort", defaultRecPort, fmt.Sprintf("port to receive on. default %d", defaultRecPort))
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(2)
	go receiver(*recPortPtr)
	go sender(*sendPortPtr)
	wg.Wait()
}

func sender(port int) {
	sender := rudp.RUDPClient{Name: "Sender", SendToPort: port, SockPort: 0}
	defer sender.Close()

	err := sender.OpenSocket()
	if err != nil {
		log.Fatalf("failed to open socket for sender. %s", err)
	}
	err = sender.BindPort()
	if err != nil {
		log.Fatalf("failed to bind port. %s", err)
	}

	// Send some bytes
	msgs := []string{
		"!",
		"hope",
		"you",
		"get",
		"all",
		"of",
		"this",
		"mi",
		"amigo",
		"¡",
	}
	for _, msg := range msgs {
		err = sender.Send([]byte(msg))
		if err != nil {
			log.Fatalf("Failed to send. %s", err)
		}
	}
}

func receiver(port int) {
	receiver := rudp.RUDPClient{Name: "Receiver", SockPort: port}
	defer receiver.Close()
	err := receiver.OpenSocket()
	if err != nil {
		log.Fatalf("failed to open socket for receiver. %s", err)
	}
	err = receiver.BindPort()
	if err != nil {
		log.Fatalf("failed to bind port. %s", err)
	}
	var resp []byte
	var n int

	// receive in a nonblocking manner forever
	for {
		resp, n, err = receiver.Receive()
		if err == nil {
			log.Printf("Received %s from 127.0.0.1:%d", resp[:n], port)
			continue
		}
		if !errors.Is(err, syscall.EAGAIN) {
			log.Fatalf("receive failed. %s", err)
		} else {
			log.Printf("Received %s. Blocking using select until it's ready and then try again.", err)
			err = receiver.WaitUntilReady()
			if err != nil {
				log.Fatalf("select failed. %s", err)
			}
		}
	}
}
