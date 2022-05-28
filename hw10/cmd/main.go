package main

import (
	pp "github.com/kevin-glare/hardcode-dev-go/hw10/pkg/pingpong"
)

func main() {
	game := pp.NewGame("Player #1", "Player #2")
	game.Start()
}
