package world

import "fmt"

type Room struct {
	Name        string
	Description string
	Players     []*Player
	Items       []*Item
}

func NewRoom(name string) *Room {
	r := Room{Name: name}
	return &r
}

func (r *Room) Print(p *Player) {
	p.Notify(r.Name + "\n")
	p.Notify(r.Description + "\n")
	p.Notify("\n")
	r.PrintPlayers(p)
	r.PrintItems(p)
}

func (r *Room) PrintPlayers(p *Player) {
	slice := r.Players
	for _, ps := range slice {
		if ps != p {
			p.Notify(fmt.Sprintf("%s is here.\n", ps.DisplayName()))
		}
	}
}

func (r *Room) PrintItems(p *Player) {
	slice := r.Items
	for _, i := range slice {
		p.Notify(i.RoomDescription + "\n")
	}
}
