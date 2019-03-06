package serve

import (
  "net/http"
  "fmt"
  "log"
  "os"

  "github.com/koeniglorenz/bwaas/pkg/store"
)

var bw store.Store
var logger *log.Logger

func New(s store.Store) http.Handler{
  logger = log.New(os.Stdout, "http: ", log.LstdFlags)
  bw = s
  r := http.NewServeMux()
  r.HandleFunc("/", handler)
  return r
}


func handler(w http.ResponseWriter, r *http.Request) {
	logger.Println(r.Method, r.RequestURI, r.UserAgent(), r.RemoteAddr)

	t := r.Header.Get("accept")

	if t == "application/json" {
		b, err := bw.GetJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error formatting JSON: %v\n", err.Error())
			return
		}
		fmt.Fprintf(w, "%s", b)
		return
	} else {
		b := bw.GetHTML()
		fmt.Fprintf(w, "%s", b)
		return
	}
}
