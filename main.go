package main

import (
	"log"
	"os"

	"go-wallet/src/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := config.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
