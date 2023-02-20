package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	helloLogger *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.helloLogger.Println("Hello world")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Unable to read!", http.StatusBadRequest)
		return
	}

	log.Println(d)
	fmt.Fprintf(rw, "Hello World! This is my data: %s", d)
}
