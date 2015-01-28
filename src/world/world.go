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
	w.Rooms = append(w.Rooms, w.DefaultRoom())
	return &w
}

func (w *World) DefaultRoom() *Room {
	r := NewRoom("The void")
	return r
}
