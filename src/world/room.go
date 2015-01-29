package world

import "fmt"

type Room struct {
	name        string
	description string
	Players     []*Player
	Items       []*Item
}

func NewRoom(name string) *Room {
	r := Room{name: name}
	return &r
}

func (r *Room) print(p *Player) {
	p.send(r.name + "\n")
	p.send(r.description + "\n")
	p.send("\n")
	r.printPlayers(p)
	r.printItems(p)
}

func (r *Room) printPlayers(p *Player) {
	slice := r.Players
	for _, ps := range slice {
		if ps != p {
			p.send(fmt.Sprintf("%s is here.\n", ps.displayName()))
		}
	}
}

func (r *Room) printItems(p *Player) {
	slice := r.Items
	for _, i := range slice {
		p.send(i.roomDescription + "\n")
	}
}
