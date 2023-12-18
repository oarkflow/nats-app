package main

import (
	"log"
	
	"github.com/nats-io/nats-server/v2/server"
)

func main() {
	// options to create nats server
	opts := &server.Options{JetStream: true}
	
	// initialize the server struct
	con, err := server.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}
	
	con.Start()
	
	log.Println("NATS server started")
	con.WaitForShutdown()
}
