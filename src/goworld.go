package main

import "github.com/codeist/goworld/src/world"

func main() {
	r := world.NewRoom("On the front step")
	i := world.NewItem("a sword", "A rusty old sword lies here stuck in the mud.", "The sword is rusty and muddy.")
	r.Items = append(r.Items, i)
	p1 := world.NewPlayer("Bilbo", r)
	world.NewPlayer("Gandalf", r)

	p1.Look()
}
