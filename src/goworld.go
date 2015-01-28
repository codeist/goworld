package main

import "github.com/codeist/goworld/src/server"

func main() {

	s := server.NewServer("127.0.0.1", 4000)
	s.Start()
}
