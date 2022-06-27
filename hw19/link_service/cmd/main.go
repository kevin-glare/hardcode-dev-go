package main

import (
	"github.com/kevin-glare/hardcode-dev-go/hw19/analytic_service/pkg/kfk"
	"github.com/kevin-glare/hardcode-dev-go/hw19/link_service/pkg/api"
	"log"
)

func main() {
	producer, err := kfk.NewProducer("localhost:9092", "storage")
	if err != nil {
		log.Fatalf(err.Error())
	}

	api.Run("localhost:8080", producer)
}
