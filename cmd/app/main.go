package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Row struct {
	Id   int
	text string
}

const portNumber = ":8060"

func home_page(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("We're live!!!!!"))
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
	//	handleRequest()
	//srv := new(wb_l0.Server)
	//if err := srv.Run(portNumber); err != nil {
	//	log.Fatal("error")
	//}

	conn, err := pgx.Connect(context.Background(), "postgres://root:password@localhost:5532/wb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	//var text string
	//err = conn.QueryRow(context.Background(), "select text from tbl").Scan(&text)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println(text)

	rows, err := conn.Query(context.Background(), "SELECT * FROM tbl")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var rowSlice []Row
	for rows.Next() {
		var r Row
		err := rows.Scan(&r.Id, &r.text)
		if err != nil {
			log.Fatal(err)
		}
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(rowSlice)

}
