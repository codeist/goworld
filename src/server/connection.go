package server

import (
	"net"

	"github.com/codeist/goworld/src/world"
)

type Connection struct {
	Player *world.Player
	Handle *net.TCPConn
}

func (c *Connection) Notify(msg string) {
	c.Handle.Write([]byte(msg + "\n"))
}
