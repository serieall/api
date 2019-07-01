package bootstrap

import (
	"github.com/nats-io/stan.go"
	"log"
)

func ackHandler(ackedNuid string, err error) {
	if err != nil {
		log.Printf("Warning: error publishing msg id %s: %v\n", ackedNuid, err.Error())
	} else {
		log.Printf("Received ack for msg id %s\n", ackedNuid)
	}
}

func InitStan() {
	sc, err := stan.Connect("serieall", "toto", stan.NatsURL("nats://localhost:4222"))

	// Simple Synchronous Publisher
	err = sc.Publish("worker_images", []byte("Hello World"))
	if err != nil {
		log.Printf("Error publishing msg %s: %v\n", err.Error())
	}

	// Close connection
	sc.Close()
}