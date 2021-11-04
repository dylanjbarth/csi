package kv

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/protobuf/proto"
)

type Router struct {
	client_listener net.Listener
	logger          *log.Logger
}

func NewRouter() *Router {
	log.Printf("Creating new router listening on port %s", ROUTER_PORT)
	l, err := net.Listen("tcp", ROUTER_PORT)
	if err != nil {
		log.Fatalf("failed to listen on socket tcp:%s %s", ROUTER_PORT, err)
	}
	return &Router{l, log.New(os.Stdout, "router ", log.LstdFlags)}
}

func (r *Router) AcceptClientConnections() {
	for {
		r.logger.Printf("Router starting to accept client connections")
		conn, err := r.client_listener.Accept()
		if err != nil {
			r.logger.Fatalf("failed to read from client: %s", err)
		}
		go r.HandleConnection(&conn)
	}
}

func (r *Router) HandleConnection(conn *net.Conn) {
	for {
		data, err := getNextMessage(conn)
		if err != nil {
			r.logger.Fatalf("failed to read from client: %s", err)
		}
		var req Request
		err = proto.Unmarshal(*data, &req)
		if err != nil {
			r.logger.Printf("unable to interpret request %s", data)
			r.Respond(conn, &Response{Code: Response_FAILURE, Message: "unable to deserialize request"})
			continue
		}
		r.logger.Printf("request: %s", req.PrettyPrint())
		switch req.Command {
		case Request_GET:
			// TODO round robin / send to follower
			res, err := r.SendToServer(&req, FOLLOWER_PORT)
			if err != nil {
				r.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("get failed: %s\n", err)})
			} else {
				r.Respond(conn, res)
			}
		case Request_SET:
			res, err := r.SendToServer(&req, LEADER_PORT)
			if err != nil {
				r.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("set failed: %s\n", err)})
			} else {
				r.Respond(conn, res)
			}
		}
	}
}

func (r *Router) SendToServer(req *Request, port string) (*Response, error) {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to leader server %s", err)
	}
	reqbytes, err := proto.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize request: %s", err)
	}
	_, err = conn.Write(*toWireFormat(&reqbytes))
	if err != nil {
		return nil, fmt.Errorf("failed to send input to server: %s", err)
	}
	data, err := getNextMessage(&conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from server: %s", err)
	}
	var resp Response
	err = proto.Unmarshal(*data, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize response from server: %s", err)
	}
	return &resp, nil

}

func (r *Router) Respond(conn *net.Conn, resp *Response) {
	r.logger.Printf("response: %s", resp.PrettyPrint())
	out, err := proto.Marshal(resp)
	if err != nil {
		r.logger.Fatalf("failed to serialize response: %s", err)
	}
	_, err = (*conn).Write(*toWireFormat(&out))
	if err != nil {
		r.logger.Fatalf("failed to respond to client: %s", err)
	}
}
