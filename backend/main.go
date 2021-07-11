package main

import (
	"log"

	"github.com/zinirun/go-music/backend/src/rest"
)

func main() {
	log.Println("Start server")
	log.Fatal(rest.RunAPI("localhost:8888"))
}
