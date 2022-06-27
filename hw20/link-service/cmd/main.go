package main

import (
	"context"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/env"
	"github.com/kevin-glare/hardcode-dev-go/hw20/link-service/pkg/api"
	"github.com/kevin-glare/hardcode-dev-go/hw20/link-service/pkg/database"
	"github.com/kevin-glare/hardcode-dev-go/hw20/link-service/pkg/repository"
	"github.com/kevin-glare/hardcode-dev-go/hw20/link-service/pkg/service"
	"log"
	"os"
)

func main() {
	env.Load("/link-service")

	mongoClient, err := database.NewClient(os.Getenv("MONGO_CONNECT_URL"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer mongoClient.Disconnect(context.Background())

	linkRepo := repository.NewLinkRepo(mongoClient.Database(os.Getenv("MONGO_DATABASE_NAME")))
	linkService := service.NewLinkService(linkRepo)

	api.Run(os.Getenv("LINK_SERVICE_HTTP_HOST"), linkService)
}
