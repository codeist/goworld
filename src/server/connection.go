package server

import (
	"net"

	"github.com/codeist/goworld/src/world"
)

type Connection struct {
	Player *world.Player
	Handle *net.TCPConn
}
