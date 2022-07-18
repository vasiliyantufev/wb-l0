package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

type Row struct {
	Id   int
	text string
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://root:password@localhost:5532/wb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(),
		"INSERT INTO tbl (json) VALUES ($1)",
		"text2")

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
