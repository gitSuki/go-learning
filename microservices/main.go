package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gitsuki/microservices/handlers"
)

func main() {
	log := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(log)
	goodbyeHandler := handlers.NewGoodbye(log)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	server := &http.Server{
		Addr:         ":8000",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	server.ListenAndServe()
}
