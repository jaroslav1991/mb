package main

import (
	"flag"
	"log"
	"mb/internal/handlers"
	"net/http"
	"strings"
)

var (
	httpRemotePort = flag.String(
		"p",
		"",
		"Remote server port, for example: -p :8080",
	)
)

func main() {
	flag.Parse()

	if strings.TrimSpace(*httpRemotePort) == "" {
		flag.Usage()
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			handlers.PutHandler(w, r)
		} else if r.Method == http.MethodGet {
			handlers.GetHandler(w, r)
		} else {
			http.Error(w, "", http.StatusMethodNotAllowed)
		}
	})

	if err := http.ListenAndServe(*httpRemotePort, nil); err != nil {
		log.Fatal(err)
	}
}
