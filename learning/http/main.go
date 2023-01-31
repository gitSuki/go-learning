package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// the response variable will return a response struct that has many fields.
	// the body field however is "io.ReadCloser" which is actually an interface!
	// why is an interface used as a type in a struct?
	// what it means is that the body field can be assigned any value that fulfils
	// the code in the io.ReadCloser interface. The io.ReadCloser is an interface
	// of interfaces. It means to be part of the io.ReadCloser interface you
	// have to meet the requirements of both of it's child interfaces.
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(response)
}
