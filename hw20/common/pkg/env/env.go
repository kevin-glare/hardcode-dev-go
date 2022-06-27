package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Load(path string) {
	err := godotenv.Load(fmt.Sprintf("%s%s/.env", currentPath(), path))
	if err != nil {
		log.Println(".env not found")
	}
}

func currentPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path
}
