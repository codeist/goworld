package world

type Item struct {
	name            string
	roomDescription string
	lookDescription string
	Room            *Room
}

func NewItem(name, roomDescription, lookDescription string) *Item {
	i := Item{name: name, roomDescription: roomDescription, lookDescription: lookDescription}
	return &i
}

func newAdmissionBell(r *Room) *Item {
	i := NewItem("a rope", "There is a rope here hanging from a bell floating in air..", "Maybe you should try pulling the rope?")
	return i
}
