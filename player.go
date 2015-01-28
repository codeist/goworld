package main

import "fmt"

type Player struct {
	Name     string
	Location *Room
}

func NewPlayer(n string, r *Room) *Player {
	p := Player{Name: n, Location: r}
	p.Move(r)
	return &p
}

func (p *Player) Move(r *Room) {
	p.Location = r
	r.Players = append(r.Players, p)
}

func (p *Player) Look() {
	p.PrintRoomPlayers(p)
	p.Location.PrintItems()
}

func (p *Player) PrintRoomPlayers(s *Player) {
	slice := p.Location.Players
	for _, ps := range slice {
		if ps != p {
			fmt.Printf("%s is here.\n", ps.Name)
		}
	}
}
