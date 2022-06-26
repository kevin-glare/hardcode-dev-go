package main

import (
	"github.com/kevin-glare/hardcode-dev-go/hw19/analytic_service/pkg/kfk"
	"github.com/kevin-glare/hardcode-dev-go/hw19/link_service/pkg/api"
	"log"
)

func main() {
	kfk, err := kfk.New([]string{"localhost:9092"}, "storage", "storage-consumer-group")
	if err != nil {
		log.Fatalf(err.Error())
	}

	api.Run("localhost:8080", kfk)
}
