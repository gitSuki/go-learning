package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	goodbyeLogger *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Goodbye!"))
}
