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

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)

	http.ListenAndServe(":8000", serveMux)
}
