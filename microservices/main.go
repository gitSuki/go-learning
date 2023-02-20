package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gitsuki/microservices/handlers"
)

func main() {
	log := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(log)
	goodbyeHandler := handlers.NewGoodbye(log)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	http.ListenAndServe(":8000", serveMux)
}
