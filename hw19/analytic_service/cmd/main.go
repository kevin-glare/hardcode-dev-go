package main

import (
	"github.com/kevin-glare/hardcode-dev-go/hw19/analytic_service/pkg/kfk"
	"log"
)

func main() {
	kfk, err := kfk.NewConsumer([]string{"localhost:9092"}, "storage", "storage-consumer-group")
	if err != nil {
		log.Fatalf(err.Error())
	}

	kfk.ConsumerRun()
}
