package main

import (
	"github.com/kevin-glare/hardcode-dev-go/hw20/api-gateway/pkg/api"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/env"
	"os"
)

func main() {
	env.Load("api-gateway")

	api.Run(os.Getenv("API_GATEWAY_HOST"))
}
