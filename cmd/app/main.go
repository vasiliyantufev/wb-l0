package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8060"

func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./web/templates/index.html")
	err := tmpl.Execute(w, "no data needed")
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}

func test_page(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test Page"))
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()
	id := values.Get("id")

	if id != "" {
		w.Write([]byte("Order " + id))
	} else {
		w.Write([]byte("Order <id>"))
	}
	w.WriteHeader(http.StatusOK)

}

func handleRequest() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", home_page)
	rtr.HandleFunc("/test", test_page)

	rtr.HandleFunc("/order", OrderHandler)

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}

func main() {
	handleRequest()
}
