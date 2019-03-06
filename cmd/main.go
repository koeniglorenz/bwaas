package main

import (
	"log"
	"net/http"
	"os"

	"github.com/koeniglorenz/bwaas/pkg/serve"
	"github.com/koeniglorenz/bwaas/pkg/store"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to buzzwords.json as a argument to the programm call.")
	}

	p := os.Args[1]

	s, err := store.New(p)
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
