package main

import (
	"log"
	"runtime"
	
	"go-nats-jetstream-example/jet"
)

func main() {
	js, err := jet.JetStreamInit()
	if err != nil {
		log.Println(err)
		return
	}
	jet.ConsumeReviews(js)
	runtime.Goexit()
}
