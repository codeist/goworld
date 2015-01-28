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

func (r *Room) Print() {
	fmt.Println(r.Name)
	r.PrintItems()
}

func (r *Room) PrintItems() {
	slice := r.Items
	for _, i := range slice {
		fmt.Println(i.RoomDescription)
	}
}
