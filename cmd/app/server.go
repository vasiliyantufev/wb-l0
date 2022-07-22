package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"github.com/vasiliyantufev/wb-l0/internal/app"
	"github.com/vasiliyantufev/wb-l0/internal/cache"
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
	w.WriteHeader(http.StatusOK)
}

func create_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./web/templates/create.html")
	err := tmpl.Execute(w, "no data needed")
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
	w.WriteHeader(http.StatusOK)
}

func test_page(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test Page"))
}

func OrderHandler(conn *pgx.Conn, cache *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		values := r.URL.Query()
		id := values.Get("id")

		cached, found := cache.GetOrder(id)

		if found == false {
			order, err := app.GetOrder(id, conn)
			if err != nil {
				log.Fatal("Order not found")
			}
			order, _ = json.Marshal(order)
			cache.PutOrder(id, string(order))
		}

		tmpl, _ := template.ParseFiles("./web/templates/show.html")
		err := tmpl.Execute(w, cached)
		if err != nil {
			log.Fatalf("execution failed: %s", err)
		}
		w.WriteHeader(http.StatusOK)
	}
}

func StoreHandler(conn *pgx.Conn, cache *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		json := r.FormValue("json")

		message, err := app.ParseMessages([]byte(json))
		if err != nil {
			log.Fatalf("Unable data: %v\n", err)
		} else {
			cache.PutOrder(message.OrderUid, json)
			err = app.InsertOrder(message, conn)
			if err != nil {
				log.Fatalf("Unable put order to database: %v\n", err)
			}
		}

		http.Redirect(w, r, "/", 301)
	}
}

func handleRequest() {

	rtr := mux.NewRouter()

	conn, err := pgx.Connect(context.Background(), "postgres://root:password@localhost:5532/wb")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	cache, err := app.GetInitialCache(conn)
	if err != nil {
		log.Fatalf("Recovery cache failed:  %v\n", err)
	}

	rtr.HandleFunc("/", home_page)
	rtr.HandleFunc("/test", test_page)

	rtr.HandleFunc("/create", create_page)
	rtr.HandleFunc("/order", OrderHandler(conn, cache))
	rtr.HandleFunc("/store", StoreHandler(conn, cache)).Methods("POST")

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}

func main() {
	handleRequest()
}
