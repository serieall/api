package bootstrap

import (
	"github.com/nats-io/stan.go"
	"log"
)

var sc stan.Conn

func InitStan() stan.Conn {
	var err error

	sc, err = stan.Connect("serieall", "api", stan.NatsURL("nats://"+GetConfig().NatsHost+":4222"))
	if err != nil {
		log.Printf("Error connecting to stan %s: %v\n", err.Error())
	}

	return sc
}

func GetStan() stan.Conn {
	return sc
}
