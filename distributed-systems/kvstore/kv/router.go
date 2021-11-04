package kv

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

type Router struct {
	client_listener net.Listener
}

func NewRouter() *Router {
	log.Printf("Creating new router listening on port %s", ROUTER_PORT)
	l, err := net.Listen("tcp", ROUTER_PORT)
	if err != nil {
		log.Fatalf("failed to listen on socket tcp:%s %s", ROUTER_PORT, err)
	}
	return &Router{l}
}

func (r *Router) AcceptClientConnections() {
	for {
		log.Printf("Router starting to accept client connections")
		conn, err := r.client_listener.Accept()
		if err != nil {
			log.Fatalf("failed to read from client: %s", err)
		}
		go r.HandleConnection(&conn)
	}
}

func (r *Router) HandleConnection(conn *net.Conn) {
	for {
		data, err := getNextMessage(conn)
		if err != nil {
			log.Fatalf("failed to read from client: %s", err)
		}
		var req Request
		err = proto.Unmarshal(*data, &req)
		if err != nil {
			log.Printf("unable to interpret request %s", data)
			r.Respond(conn, &Response{Code: Response_FAILURE, Message: "unable to deserialize request"})
			continue
		}
		log.Printf("request: %s", req.PrettyPrint())
		switch req.Command {
		case Request_GET:
			// TODO round robin / send to follower
			res, err := r.SendToServer(&req)
			if err != nil {
				r.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("get failed: %s\n", err)})
			} else {
				r.Respond(conn, &Response{Code: Response_SUCCESS, Message: fmt.Sprintf("%s\n", res)})
			}
		case Request_SET:
			_, err = r.SendToServer(&req)
			if err != nil {
				r.Respond(conn, &Response{Code: Response_FAILURE, Message: fmt.Sprintf("set failed: %s\n", err)})
			} else {
				r.Respond(conn, &Response{Code: Response_SUCCESS, Message: fmt.Sprintf("Success: %s=%s\n", req.Item.Key, req.Item.Value)})
			}
		}
	}
}

func (r *Router) SendToServer(req *Request) (*Response, error) {
	conn, err := net.Dial("tcp", LEADER_PORT)
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
