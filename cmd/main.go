package main

import (
	"log"
	"net/http"
	"flag"

	"github.com/koeniglorenz/bwaas/pkg/serve"
	"github.com/koeniglorenz/bwaas/pkg/store"
)

func main() {
	buzzwords := flag.String("buzzwords", "buzzwords.json", "path to JSON-file with buzzwords")

	flag.Parse()

	s, err := store.New(*buzzwords)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := serve.New(s)

	log.Println("Starting HTTP-Server at :8080...")
	err = http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Fatalf("Error starting up HTTP-Server: %v", err)
	}
}
