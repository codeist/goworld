package world

import (
	"fmt"
	"net"
)

type Player struct {
	Connection *net.TCPConn
	Name       string
	Room       *Room
}

func NewPlayer(n string, r *Room) *Player {
	p := Player{Name: n, Room: r}
	p.Move(r)
	return &p
}

func (p *Player) Move(r *Room) {
	p.Room = r
	r.Players = append(r.Players, p)
}

func (p *Player) Look() {
	p.PrintRoomPlayers()
	p.Room.PrintItems()
}

func (p *Player) Notify(m string) {
	p.Connection.Write([]byte(m))
}

func (p *Player) DisplayName() string {
	if len(p.Name) > 0 {
		return p.Name
	} else {
		return "A cloaked figure"
	}
}

func (p *Player) PrintRoomPlayers() {
	slice := p.Room.Players
	for _, ps := range slice {
		if ps != p {
			p.Notify(fmt.Sprintf("%s is here.\n", ps.DisplayName()))
		}
	}
}
