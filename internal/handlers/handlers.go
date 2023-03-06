package handlers

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var queues map[string]*Queue = map[string]*Queue{}

type Queue struct {
	Message []string
	Waiting sync.Mutex
}

func (q *Queue) Add(v string) {
	q.Message = append(q.Message, v)
}

func (q *Queue) Get() (string, bool) {
	q.Waiting.Lock()
	defer q.Waiting.Unlock()
	if len(q.Message) > 0 {
		value := q.Message[0]

		q.Message = q.Message[1:]

		return value, true
	}

	return "", false
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
	} else {
		w.WriteHeader(http.StatusBadRequest)
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

	timeValue := r.Form.Get("timeout")
	timeout := time.Nanosecond

	sleepTime := time.Duration(0)

	if timeValue != "" {
		timeoutParams, err := time.ParseDuration(timeValue + "s")
		if err == nil {
			timeout = timeoutParams
			sleepTime = time.Second
		} else {
			log.Println(err)
			return
		}
	}

	for i := time.Duration(0); i < timeout; i += sleepTime {
		value, ok := queues[queue].Get()
		if ok {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, value)
			return
		}
		time.Sleep(sleepTime)
	}

	w.WriteHeader(http.StatusNotFound)
}
