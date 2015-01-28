package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/codeist/goworld/src/world"
)

type Server struct {
	Host  string
	Port  uint
	World *world.World
}

func exitIf(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func NewServer(host string, port uint) *Server {
	s := Server{Host: host, Port: port, World: world.NewWorld()}
	return &s
}

func (s *Server) Start() {
	a, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(s.Host, fmt.Sprintf("%d", s.Port)))
	exitIf(err)

	l, err := net.ListenTCP("tcp", a)
	exitIf(err)

	log.Printf("The server started and listening on port %d.\n", s.Port)

	for {
		c, err := l.AcceptTCP()
		exitIf(err)
		s.HandleIncoming(c)
	}
}

func (s *Server) Notify(m string) {
	log.Println(m)
}

func (s *Server) NotifyIncoming(c *net.TCPConn) {
	s.Notify(fmt.Sprintf("A player has connected from %s", c.RemoteAddr()))
	s.Notify(fmt.Sprintf("There are now %d active connections", len(s.World.Players)))
}

func (s *Server) HandleIncoming(conn *net.TCPConn) *world.Player {
	p := world.Player{Connection: conn}
	s.World.Players = append(s.World.Players, &p)
	s.NotifyIncoming(conn)
	p.Move(s.World.DefaultRoom())
	p.Look()
	return &p
}
