package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var queues map[string]*Queue = map[string]*Queue{}

type Queue struct {
	Message []string
}

func (q *Queue) Add(v string) {
	q.Message = append(q.Message, v)
}

func (q *Queue) Get() string {
	value := q.Message[0]

	q.Message = q.Message[1:]

	return value
}

func PutHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println("can't parse put form", err)
		return
	}
	value := r.Form.Get("v")
	queue := r.URL.Path[1:]

	if _, exist := queues[queue]; !exist {
		queues[queue] = &Queue{}
	}

	if value != "" {
		queues[queue].Add(value)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "not empty")
		fmt.Println(queues)
		fmt.Println(queues[queue])
	} else {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "empty")
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("can't parse get form", err)
		return
	}
	queue := r.URL.Path[1:]

	if _, exist := queues[queue]; !exist {
		queues[queue] = &Queue{}
	}

	if len(queues[queue].Message) > 0 {
		fmt.Println(queues[queue].Get())
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "not nil")

		fmt.Println(queues[queue].Message)
	} else {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "empty body")
	}
}
