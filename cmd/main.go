package main

import (
	"flag"
	"log"

	"github.com/koeniglorenz/bwaas/pkg/serve"
	"github.com/koeniglorenz/bwaas/pkg/store"
)

func main() {
	buzzwords := flag.String("buzzwords", "buzzwords.json", "path to JSON-file with buzzwords")

	flag.Parse()

	store, err := store.New(*buzzwords)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := serve.New(store)
	server.Start()
}
