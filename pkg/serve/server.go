package serve

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/koeniglorenz/bwaas/pkg/store"
)

type server struct {
	logger *log.Logger
	store store.Store
	mux *http.ServeMux
}

func New(store store.Store) *server {
	s := &server{}
	s.logger = log.New(os.Stdout, "http: ", log.LstdFlags)
	s.store = store
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/", s.handler)
	return s
}

func (s server) handler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println(r.Method, r.RequestURI, r.UserAgent(), r.RemoteAddr)

	t := r.Header.Get("accept")

	if t == "application/json" {
		b, err := s.store.GetJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error formatting JSON: %v\n", err.Error())
			return
		}
		fmt.Fprintf(w, "%s", b)
		return
	} else {
		b := s.store.GetHTML()
		fmt.Fprintf(w, "%s", b)
		return
	}
}

func (s server) Start() error {
	log.Println("Starting HTTP-Server at :8080...")
	err := http.ListenAndServe("0.0.0.0:8080", s.mux)
	return err
}
