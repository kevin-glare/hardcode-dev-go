package main

import (
	"github.com/kevin-glare/hardcode-dev-go/hw20/analytic-service/pkg/api"
	"github.com/kevin-glare/hardcode-dev-go/hw20/analytic-service/pkg/service"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/env"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/kfk"
	"log"
	"os"
)

func main() {
	env.Load("analytic-service")

	stat := service.NewStatistic()

	consumer, err := kfk.NewConsumer([]string{os.Getenv("KAFKA_BROKER")}, os.Getenv("KAFKA_TOPIC"), os.Getenv("KAFKA_GROUP_ID"), stat.Update)
	if err != nil {
		log.Fatalf(err.Error())
	}

	go consumer.Run()
	api.Run(os.Getenv("ANALYTIC_SERVICE_HTTP_HOST"), stat)
}
