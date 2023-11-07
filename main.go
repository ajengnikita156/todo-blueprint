package main

import (
	"log"

	"membuattasktodo/routes"
)

func main() {
	err := routes.Init()
	if err != nil {
		log.Fatalf("Error start the server with err: %s", err)
	}
}
