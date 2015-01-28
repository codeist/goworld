package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/codeist/goworld/src/world"
)

type Server struct {
	Host        string
	Port        uint
	Connections []*Connection
}

func exitIf(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func NewServer(host string, port uint) *Server {
	s := Server{Host: host, Port: port}
	return &s
}

func (s *Server) Start() {
	a, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(s.Host, fmt.Sprintf("%d", s.Port)))
	exitIf(err)

	l, err := net.ListenTCP("tcp", a)
	exitIf(err)

	log.Printf("The server started and listening on port %d.\n", s.Port)

	for {
		conn, err := l.AcceptTCP()
		exitIf(err)
		s.NewConnection(conn)
	}
}

func (s *Server) NewConnection(conn *net.TCPConn) *Connection {
	c := Connection{Player: &world.Player{}, Handle: conn}
	s.Connections = append(s.Connections, &c)
	log.Printf("A player has connected from %s", c.Handle.RemoteAddr())
	log.Printf("There are now %d active connections", len(s.Connections))
	return &c
}
