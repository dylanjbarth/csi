package kv

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/protobuf/proto"
)

type ConnHandler interface {
	Log(msg string)
	GetListener() net.Listener
	HandleConnection(*net.Conn)
}

type Server struct {
	request_listener net.Listener
	s                *Storage
	logger           *log.Logger
}

type Leader struct {
	Server
}

type Follower struct {
	Server
}

func NewLeader(path string) *Leader {
	log.Printf("Creating new leader listening on port %s and storing data %s", LEADER_PORT, path)
	return &Leader{Server{initListener(LEADER_PORT), NewStorage(path, false), log.New(os.Stdout, "leader: ", log.LstdFlags)}}
}

func NewFollower(path string) *Follower {
	log.Printf("Creating new follower listening on port %s", FOLLOWER_PORT)
	return &Follower{Server{initListener(FOLLOWER_PORT), NewStorage(path, false), log.New(os.Stdout, "follower: ", log.LstdFlags)}}
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

func (l *Leader) Log(msg string) {
	l.logger.Print(msg)
}

func (f *Follower) Log(msg string) {
	f.logger.Print(msg)
}

func AcceptRouterConnections(n ConnHandler) {
	for {
		n.Log("Starting to accept connections ")
		conn, err := n.GetListener().Accept()
		if err != nil {
			n.Log(fmt.Sprintf("failed to connect: %s", err))
		}
		go n.HandleConnection(&conn)
	}
}

// Leader can handle writes and reads, also responsible for replicating to all followers.
func (l *Leader) HandleConnection(conn *net.Conn) {
	for {
		data, err := getNextMessage(conn)
		if err != nil {
			log.Fatalf("leader failed to read from connection: %s", err)
		}
		var req Request
		err = proto.Unmarshal(*data, &req)
		if err != nil {
			l.Log(fmt.Sprintf("unable to interpret request %s", data))
			l.Respond(conn, &Response{Code: Response_FAILURE, Message: "unable to deserialize request"})
			continue
		}
		l.Log(fmt.Sprintf("request: %s", req.PrettyPrint()))
		switch req.Command {
		case Request_GET:
			// TODO leader can obviously handle get requests too, but enforcing this for now.
			l.Respond(conn, &Response{Code: Response_FAILURE, Message: "leader only handles writes\n"})
			// res, err := s.s.Get(req.Item.Key)
			// if err != nil {
			// 	l.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("get failed: %s\n", err)})
			// } else {
			// 	l.Respond(conn, &Response{Code: Response_SUCCESS, Message: fmt.Sprintf("%s\n", res)})
			// }
		case Request_SET:
			err = l.s.Set(req.Item.Key, req.Item.Value)
			if err != nil {
				l.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("set failed: %s\n", err)})
			} else {
				l.Respond(conn, &Response{Code: Response_SUCCESS, Message: fmt.Sprintf("Success: %s=%s\n", req.Item.Key, req.Item.Value)})
			}
			// async replicate to followers (ignoring retries for now!)

		}
	}
}

// Follower can handle reads from clients and must also handle writes from the leader.
func (f *Follower) HandleConnection(conn *net.Conn) {
	for {
		data, err := getNextMessage(conn)
		if err != nil {
			log.Fatalf("follower failed to read from connection: %s", err)
		}
		var req Request
		err = proto.Unmarshal(*data, &req)
		if err != nil {
			f.Log(fmt.Sprintf("unable to interpret request %s", data))
			f.Respond(conn, &Response{Code: Response_FAILURE, Message: "unable to deserialize request"})
			continue
		}
		f.Log(fmt.Sprintf("request: %s", req.PrettyPrint()))
		switch req.Command {
		case Request_GET:
			res, err := f.s.Get(req.Item.Key)
			if err != nil {
				f.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("get failed: %s\n", err)})
			} else {
				f.Respond(conn, &Response{Code: Response_SUCCESS, Message: fmt.Sprintf("%s\n", res)})
			}
		case Request_SET:
			f.Respond(conn, &Response{Code: Response_FAILURE, Message: "follower only handles reads\n"})
		}
	}
}

func (s *Server) Respond(conn *net.Conn, resp *Response) {
	s.logger.Printf(fmt.Sprintf("response: %s", resp.PrettyPrint()))
	out, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalf("failed to serialize response: %s", err)
	}
	_, err = (*conn).Write(*toWireFormat(&out))
	if err != nil {
		log.Fatalf("failed to respond to client: %s", err)
	}
}
