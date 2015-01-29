package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/codeist/goworld/src/world"
)

type Server struct {
	host        string
	port        uint
	Connections []*net.TCPConn
	World       *world.World
}

func exitIf(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func NewServer(host string, port uint) *Server {
	s := Server{host: host, port: port, World: world.NewWorld()}
	return &s
}

func (s *Server) Start() {
	a, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(s.host, fmt.Sprintf("%d", s.port)))
	exitIf(err)

	l, err := net.ListenTCP("tcp", a)
	exitIf(err)

	log.Printf("The server started and listening on port %d.\n", s.port)

	for {
		c, err := l.AcceptTCP()
		exitIf(err)
		s.handleIncoming(c)
	}
}

func (s *Server) send(m string) {
	log.Println(m)
}

func (s *Server) notifyIncoming(c *net.TCPConn) {
	s.send(fmt.Sprintf("A player has connected from %s", c.RemoteAddr()))
	s.send(fmt.Sprintf("There are now %d active connections", len(s.World.Players)))
}

func (s *Server) handleIncoming(conn *net.TCPConn) *world.Player {
	p := world.Player{Connection: conn}
	s.Connections = append(s.Connections, conn)
	p.EnterWorld(s.World)
	s.notifyIncoming(conn)
	return &p
}
