package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Unable to read!", http.StatusBadRequest)
			return
		}

		log.Println(d)
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye world")
	})
	http.ListenAndServe(":8000", nil)
}
