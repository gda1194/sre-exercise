package main

import (
	"log"
	"sre-exercise/cmd/api/bootstrap"
)

func main() {
	err := bootstrap.Run()
	// Hola
	if err != nil {
		log.Fatal(err)
	}
}
