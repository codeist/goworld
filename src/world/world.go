package world

import (
	"log"
	"os"
)

type World struct {
	p []*Player
	r []*Room
}

func exitIf(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
