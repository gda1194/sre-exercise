package main

import (
	"log"
	"sre-exercise/cmd/api/bootstrap"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatal(err)
	}
}
