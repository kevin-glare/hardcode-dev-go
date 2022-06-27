package main

import (
	"github.com/kevin-glare/hardcode-dev-go/hw20/api-gateway/pkg/api"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/env"
	"os"
)

func main() {
	env.Load("api-gateway")
	env.Load("analytic-service")
	env.Load("cache-service")
	env.Load("link-service")

	api.Run(os.Getenv("API_GATEWAY_HOST"))
}
