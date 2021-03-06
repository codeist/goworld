package world

import (
	"log"
	"os"
)

type World struct {
	Players []*Player
	Rooms   []*Room
}

func exitIf(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func NewWorld() *World {
	w := World{}
	w.Rooms = append(w.Rooms, w.defaultRoom())
	return &w
}

func (w *World) defaultRoom() *Room {
	if len(w.Rooms) == 0 {
		r := NewRoom("The void")
		i := newAdmissionBell(r)
		r.Items = append(r.Items, i)
		return r
	} else {
		return w.Rooms[0]
	}
}
