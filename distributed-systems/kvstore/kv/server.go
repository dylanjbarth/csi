package kv

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/protobuf/proto"
)

type Server struct {
	l net.Listener
	s *Storage
}

func NewServer(path string) *Server {
	l := initListener()
	return &Server{l, NewStorage(path, false)}
}

func initListener() net.Listener {
	// Unlink socket to listen (ignoring errors)
	os.Remove(SOCKET_FILE)
	l, err := net.Listen("unix", SOCKET_FILE)
	if err != nil {
		log.Fatalf("failed to listen on socket %s %s", SOCKET_FILE, err)
	}
	return l
}

func (s *Server) AcceptConnections() {
	for {
		log.Printf("Starting to accept connections")
		conn, err := s.l.Accept()
		if err != nil {
			log.Fatalf("failed to read from client: %s", err)
		}
		go s.HandleConnection(&conn)
	}
}

func (s *Server) HandleConnection(conn *net.Conn) {
	for {
		data, err := getNextMessage(conn)
		if err != nil {
			log.Fatalf("failed to read from client: %s", err)
		}
		log.Printf("Got %s", data)
		var req Request
		err = proto.Unmarshal(*data, &req)
		if err != nil {
			s.Respond(conn, &Response{Code: Response_FAILURE, Message: "unable to deserialize request"})
			continue
		}

		switch req.Command {
		case Request_GET:
			res, err := s.s.Get(req.Item.Key)
			if err != nil {
				s.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("get failed: %s\n", err)})
			} else {
				s.Respond(conn, &Response{Code: Response_SUCCESS, Message: fmt.Sprintf("%s\n", res)})
			}
		case Request_SET:
			err = s.s.Set(req.Item.Key, req.Item.Value)
			if err != nil {
				s.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("set failed: %s\n", err)})
			} else {
				s.Respond(conn, &Response{Code: Response_SUCCESS, Message: fmt.Sprintf("Success: %s=%s\n", req.Item.Key, req.Item.Value)})
			}
		}
	}
}

func (s *Server) Respond(conn *net.Conn, resp *Response) {
	out, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalf("failed to serialize response: %s", err)
	}
	_, err = (*conn).Write(*toWireFormat(&out))
	if err != nil {
		log.Fatalf("failed to respond to client: %s", err)
	}
}
