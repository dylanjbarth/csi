package kv

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Server struct {
	l net.Listener
	p *Parser
	s *Storage
}

func NewServer(path string) *Server {
	l := initListener()
	return &Server{l, NewParser(), NewStorage(path, false)}
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
		go s.HandleConnection(conn)
	}
}

func (s *Server) HandleConnection(conn net.Conn) {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read from client: %s", err)
		}
		log.Printf("Got %s", data)
		cmdgrp, err := s.p.Parse(data)
		if err != nil {
			s.Respond(conn, fmt.Sprintf("parsing failed: %s", err))
		}
		log.Printf("Parsed %s", cmdgrp)
		switch cmdgrp.Cmd {
		case CMD_GET:
			res, err := s.s.Get(cmdgrp.Args[0])
			if err != nil {
				s.Respond(conn, fmt.Sprintf("get failed: %s", err))
			} else {
				s.Respond(conn, fmt.Sprintf("%s\n", res))
			}
		case CMD_SET:
			err = s.s.Set(cmdgrp.Args[0], cmdgrp.Args[1])
			if err != nil {
				s.Respond(conn, fmt.Sprintf("set failed: %s", err))
			} else {
				s.Respond(conn, "\n")
			}
		}
	}
}

func (s *Server) Respond(conn net.Conn, resp string) {
	_, err := conn.Write([]byte(resp))
	if err != nil {
		log.Fatalf("failed to respond to client: %s", err)
	}
}
