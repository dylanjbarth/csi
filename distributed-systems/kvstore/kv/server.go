package kv

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

type ConnHandler interface {
	GetListener() net.Listener
	HandleConnection(*net.Conn)
}

type Server struct {
	request_listener net.Listener
	s                *Storage
}

type Leader struct {
	Server
}

type Follower struct {
	Server
}

func NewLeader(path string) *Leader {
	log.Printf("Creating new leader listening on port %s and storing data %s", LEADER_PORT, path)
	return &Leader{Server{initListener(LEADER_PORT), NewStorage(path, false)}}
}

func NewFollower(path string) *Follower {
	log.Printf("Creating new follower listening on port %s", FOLLOWER_PORT)
	return &Follower{Server{initListener(FOLLOWER_PORT), NewStorage(path, false)}}
}

func initListener(port string) net.Listener {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s %s", port, err)
	}
	return l
}

func (l *Leader) GetListener() net.Listener {
	return l.request_listener
}

func (f *Follower) GetListener() net.Listener {
	return f.request_listener
}

func AcceptRouterConnections(n ConnHandler) {
	for {
		log.Printf("Starting to accept connections ")
		conn, err := n.GetListener().Accept()
		if err != nil {
			log.Fatalf("failed to connect: %s", err)
		}
		go n.HandleConnection(&conn)
	}
}

// Leader can handle writes and reads, also responsible for replicating to all followers.
func (s *Leader) HandleConnection(conn *net.Conn) {
	for {
		data, err := getNextMessage(conn)
		if err != nil {
			log.Fatalf("failed to read from connection: %s", err)
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
			// TODO leader can obviously handle get requests too, but enforcing this for now.
			s.Respond(conn, &Response{Code: Response_FAILURE, Message: "leader only handles writes\n"})
			// res, err := s.s.Get(req.Item.Key)
			// if err != nil {
			// 	s.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("get failed: %s\n", err)})
			// } else {
			// 	s.Respond(conn, &Response{Code: Response_SUCCESS, Message: fmt.Sprintf("%s\n", res)})
			// }
		case Request_SET:
			err = s.s.Set(req.Item.Key, req.Item.Value)
			if err != nil {
				s.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("set failed: %s\n", err)})
			} else {
				s.Respond(conn, &Response{Code: Response_SUCCESS, Message: fmt.Sprintf("Success: %s=%s\n", req.Item.Key, req.Item.Value)})
			}
			// async replicate to followers (ignoring retries for now!)

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
