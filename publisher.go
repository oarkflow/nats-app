package main

import (
	"log"
	
	"go-nats-jetstream-example/jet"
)

func main() {
	js, err := jet.JetStreamInit()
	if err != nil {
		log.Println(err)
		return
	}
	jet.PublishReviews(js)
}
