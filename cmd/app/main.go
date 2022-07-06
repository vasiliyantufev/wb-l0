package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("We're live!!!!!"))
}

func main() {

	http.HandleFunc("/", HomePageHandler)
	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
