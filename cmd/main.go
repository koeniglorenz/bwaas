package main

import (
	"flag"
	"log"

	"github.com/koeniglorenz/bwaas/pkg/serve"
	"github.com/koeniglorenz/bwaas/pkg/store"
)

func main() {
	buzzwords := flag.String("buzzwords", "buzzwords.json", "path to JSON-file with buzzwords")
	port := flag.Int("port", 8080, "port for the server to listen on")

	flag.Parse()

	store, err := store.New(*buzzwords)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := serve.New(store, *port)
	server.Start()
}
