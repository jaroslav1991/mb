package main

import (
	"log"
	"mb/internal/handlers"
	"net/http"
)

func main() {
	http.Handle("/", handlers.HandlerCommon())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
