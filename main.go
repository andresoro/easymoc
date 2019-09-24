package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/andresoro/easymock/server"
)

func main() {
	s, err := server.New("./client/build", "./db")
	if err != nil {
		log.Fatalf("server failed to init with err: %e", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			s.ShutDown()
			os.Exit(0)
		}
	}()

	s.Run()
}
