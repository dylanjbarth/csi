package main

import (
	"flag"
	"fmt"
	"log"

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

const defaultSendPort, defaultRecPort = 8000, 8001

func main() {
	sendPortPtr := flag.Int("sendPort", defaultSendPort, fmt.Sprintf("port to send on. default %d", defaultSendPort))
	recPortPtr := flag.Int("recPort", defaultRecPort, fmt.Sprintf("port to receive on. default %d", defaultRecPort))
	flag.Parse()

	receiver := rudp.RUDPClient{Name: "Receiver", SockPort: *recPortPtr}
	sender := rudp.RUDPClient{Name: "Sender", SendToPort: *sendPortPtr, SockPort: 0}
	defer receiver.Close()
	defer sender.Close()

	err := receiver.OpenSocket()
	if err != nil {
		log.Fatalf("failed to open socket for receiver. %s", err)
	}
	err = sender.OpenSocket()
	if err != nil {
		log.Fatalf("failed to open socket for sender. %s", err)
	}
	err = receiver.BindPort()
	if err != nil {
		log.Fatalf("failed to bind port. %s", err)
	}
	err = sender.BindPort()
	if err != nil {
		log.Fatalf("failed to bind port. %s", err)
	}

	// Send some bytes
	msgs := []string{
		"hope",
		"you",
		"get",
		"all",
		"of",
		"this",
	}
	for _, msg := range msgs {
		err = sender.Send([]byte(msg))
		if err != nil {
			log.Fatalf("Failed to send. %s", err)
		}
		resp, n, err := receiver.Receive()
		if err != nil {
			log.Fatalf("receive failed. %s", err)
		}
		log.Printf("Received %s from 127.0.0.1:%d", resp[:n], *recPortPtr)
	}
}
