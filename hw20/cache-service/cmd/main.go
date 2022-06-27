package main

import (
	"github.com/kevin-glare/hardcode-dev-go/hw20/cache-service/pkg/api"
	"github.com/kevin-glare/hardcode-dev-go/hw20/cache-service/pkg/cache"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/env"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/kfk"
	"log"
	"os"
)

func main() {
	env.Load("cache-service")

	redis := cache.NewCache(os.Getenv("REDIS_URL"))

	consumer, err := kfk.NewConsumer([]string{os.Getenv("KAFKA_BROKER")}, os.Getenv("KAFKA_TOPIC"), os.Getenv("KAFKA_GROUP_ID"), redis.Set)
	if err != nil {
		log.Fatalf(err.Error())
	}

	go consumer.ConsumerRun()

	api.Run(os.Getenv("CACHE_SERVICE_HTTP_HOST"), redis, consumer)
}
