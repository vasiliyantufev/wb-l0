package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

type Row struct {
	Id   int
	text string
}

const portNumber = ":8060"

func home_page(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("We're live!!!!!"))
	tmpl, _ := template.ParseFiles("./web/templates/index.html")
	err := tmpl.Execute(w, "no data needed")
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}

func main() {
	handleRequest()

}
