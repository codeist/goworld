package world

import "net"

type Player struct {
	Connection *net.TCPConn
	name       string
	Room       *Room
}

func NewPlayer(n string, r *Room) *Player {
	p := Player{name: n, Room: r}
	p.move(r)
	return &p
}

func (p *Player) move(r *Room) {
	p.Room = r
	r.Players = append(r.Players, p)
	p.look()
}

func (p *Player) look() {
	p.Room.print(p)
}

func (p *Player) send(m string) {
	p.Connection.Write([]byte(m))
}

func (p *Player) displayName() string {
	if len(p.name) > 0 {
		return p.name
	} else {
		return "A cloaked figure"
	}
}

func (p *Player) EnterWorld(w *World) {
	w.Players = append(w.Players, p)
	p.move(w.defaultRoom())
}
