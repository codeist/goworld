package main

func main() {
	r := NewRoom("On the front step")
	i := NewItem("a sword", "A rusty old sword lies here stuck in the mud.", "The sword is rusty and muddy.")
	r.Items = append(r.Items, i)
	p1 := NewPlayer("Bilbo", r)
	NewPlayer("Gandalf", r)

	p1.Look()
}
