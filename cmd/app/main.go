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

//func test_page(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Test Page")
//}
//func handler(w http.ResponseWriter, r *http.Request, mystr string) {
//	println(mystr)
//}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Order: %v\n", vars["id"])
}

func handleRequest() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", home_page)

	//hi := "hello"
	//rtr.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
	//	handler(w, r, hi)
	//})

	rtr.HandleFunc("/order/{id}/", OrderHandler)

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}

func main() {
	handleRequest()
}
