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
