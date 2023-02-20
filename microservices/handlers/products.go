package handlers

import (
	"log"
	"net/http"

	"github.com/gitsuki/microservices/data"
)

type Products struct {
	productLogger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	productsList := data.GetProducts()
	err := productsList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}
