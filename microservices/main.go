package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gitsuki/microservices/handlers"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":8000", "Bind address for the server")

func main() {
	env.Parse()

	customLog := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create handlers and register them to servemux
	productHandler := handlers.NewProducts(customLog)
	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	server := &http.Server{
		Addr:         *bindAddress,
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			customLog.Fatal(err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// block until a signal is received.
	sig := <-signalChannel
	customLog.Println("Recieved terminate, graceful shutdown. Got signal: ", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	timeoutCTX, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutCTX)
}
