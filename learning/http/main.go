package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// the response variable will return a response struct that has many fields.
	// the body field however is "io.ReadCloser" which is actually an interface!
	// why is an interface used as a type in a struct?
	// what it means is that the body field can be assigned any value that fulfils
	// the code in the io.ReadCloser interface. The io.ReadCloser is an interface
	// of interfaces. It means to be part of the io.ReadCloser interface you
	// have to meet the requirements of both of it's child interfaces.

	// the reader interface is useful, think of it sorta like the fmt.println function
	// it can take multiple different data types like strings, arrays, slices, etc.
	// the reader does something similar by translating data into []byte which
	// is much easier for other functions to work with.
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, response.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Wrote %d", len(bs))
	return len(bs), nil
}
