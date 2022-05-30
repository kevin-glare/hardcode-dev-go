package main

import (
	"bufio"
	"github.com/kevin-glare/hardcode-dev-go/hw11/pkg/netsrv"
	"log"
	"os"
)

func main() {
	var query string
	var err error

	client, err  := netsrv.NewClient()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer client.Close()

	r := bufio.NewReader(os.Stdin)

	for {
		query, err = r.ReadString('\n')
		if err != nil {
			log.Println(err.Error())
		}

		log.Println(client.SendQuery(query))
	}
}