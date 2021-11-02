package kv

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/protobuf/proto"
)

type ConnHandler interface {
	GetListener() net.Listener
	HandleConnection(*net.Conn)
}

type Server struct {
	l net.Listener
	s *Storage
}

type Leader struct {
	Server
}

type Follower struct {
	Server
}

func NewLeader(path string) *Leader {
	l := initListener()
	return &Leader{Server{l, NewStorage(path, false)}}
}

func NewFollower(path string) *Follower {
	l := initListener()
	return &Follower{Server{l, NewStorage(path, false)}}
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

func (l *Leader) GetListener() net.Listener {
	return l.l
}

func (f *Follower) GetListener() net.Listener {
	return f.l
}

func AcceptConnections(n ConnHandler) {
	for {
		log.Printf("Starting to accept connections ")
		conn, err := n.GetListener().Accept()
		if err != nil {
			log.Fatalf("failed to read from client: %s", err)
		}
		go n.HandleConnection(&conn)
	}
}

// Leader can handle writes and reads, also responsible for replicating to all followers.
func (s *Leader) HandleConnection(conn *net.Conn) {
	for {
		data, err := getNextMessage(conn)
		if err != nil {
			log.Fatalf("failed to read from client: %s", err)
		}
		var req Request
		err = proto.Unmarshal(*data, &req)
		if err != nil {
			log.Printf("unable to interpret request %s", data)
			s.Respond(conn, &Response{Code: Response_FAILURE, Message: "unable to deserialize request"})
			continue
		}
		log.Printf("request: %s", req.PrettyPrint())
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

// Follower can handle reads from clients and must also handle writes from the leader.
func (s *Follower) HandleConnection(conn *net.Conn) {
	for {
		data, err := getNextMessage(conn)
		if err != nil {
			log.Fatalf("failed to read from client: %s", err)
		}
		var req Request
		err = proto.Unmarshal(*data, &req)
		if err != nil {
			log.Printf("unable to interpret request %s", data)
			s.Respond(conn, &Response{Code: Response_FAILURE, Message: "unable to deserialize request"})
			continue
		}
		log.Printf("request: %s", req.PrettyPrint())
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
				s.Respond(conn, &Response{Code: Response_SUCCESS, Message: fmt.Sprintf("success: %s=%s\n", req.Item.Key, req.Item.Value)})
			}
		}
	}
}

func (s *Server) Respond(conn *net.Conn, resp *Response) {
	log.Printf("response: %s", resp.PrettyPrint())
	out, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalf("failed to serialize response: %s", err)
	}
	_, err = (*conn).Write(*toWireFormat(&out))
	if err != nil {
		log.Fatalf("failed to respond to client: %s", err)
	}
}
