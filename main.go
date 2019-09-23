package main

import (
	"log"

	"github.com/andresoro/easymock/server"
)

func main() {
	s, err := server.New("./client/build", "./db")
	if err != nil {
		log.Fatalf("server failed to init with err: %e", err)
	}

	s.Run()
}
