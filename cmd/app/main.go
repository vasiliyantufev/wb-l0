package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"github.com/vasiliyantufev/wb-l0/internal/app"
	"github.com/vasiliyantufev/wb-l0/internal/models"
	"html/template"
	"log"
	"net/http"
	"os"
)

const portNumber = ":8060"

func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./web/templates/index.html")
	err := tmpl.Execute(w, "no data needed")
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
	w.WriteHeader(http.StatusOK)
}

func test_page(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test Page"))
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()
	id := values.Get("id")

	conn, err := pgx.Connect(context.Background(), "postgres://root:password@localhost:5532/wb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var order models.Order

	obj := app.GetOrder(id, conn)
	err = json.Unmarshal(obj, &order)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, _ := template.ParseFiles("./web/templates/show.html")
	errr := tmpl.Execute(w, order)
	if errr != nil {
		log.Fatalf("execution failed: %s", errr)
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
