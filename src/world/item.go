package world

type Item struct {
	Name            string
	RoomDescription string
	LookDescription string
	Location        *Room
}

func NewItem(name, roomDescription, lookDescription string) *Item {
	i := Item{Name: name, RoomDescription: roomDescription, LookDescription: lookDescription}
	return &i
}

func NewAdmissionBell(r *Room) *Item {
	i := NewItem("a rope", "There is a rope here hanging from a bell floating in air..", "Maybe you should try pulling the rope?")
	return i
}
