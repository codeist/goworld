package world

import "net"

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
	p.Look()
}

func (p *Player) Look() {
	p.Room.Print(p)
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
