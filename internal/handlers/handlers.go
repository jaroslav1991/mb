package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var arr []string

func HandlerCommon() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Println("can't parse put form", err)
			return
		}
		queue := r.Form.Get("queue")

		if r.Method == http.MethodPut {
			if queue != "" {
				arr = append(arr, queue)
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, "not empty")
				fmt.Println(arr)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, "empty")
			}
		}
		if r.Method == http.MethodGet {
			for i := 0; i < len(arr)+1; i++ {
				if len(arr) > 0 {
					w.WriteHeader(http.StatusOK)
					io.WriteString(w, "not nil")
					fmt.Println(arr[i])
					arr = remove(arr, i)
					fmt.Println(arr)
					fmt.Println(len(arr))
					break
				}
				w.WriteHeader(http.StatusNotFound)
				io.WriteString(w, "empty body")
			}
		}

	}
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
