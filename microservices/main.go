package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gitsuki/microservices/handlers"
)

func main() {
	customLog := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(customLog)
	goodbyeHandler := handlers.NewGoodbye(customLog)

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

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			customLog.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	customLog.Println("Recieved terminate, graceful shutdown", sig)

	timeoutCTX, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutCTX)
}
