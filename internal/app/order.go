package app

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"github.com/vasiliyantufev/wb-l0/internal/cache"
	"github.com/vasiliyantufev/wb-l0/internal/models"
	"log"
	"time"
)

func InsertOrder(message *models.Order, conn *pgx.Conn) (err error) {

	log.Println(message)

	jsonObj, _ := json.Marshal(message)
	_, err = conn.Exec(context.Background(),
		"INSERT INTO orders VALUES ($1, $2)", message.OrderUid, jsonObj)
	if err != nil {
		return err
	}
	return
}

func GetOrder(id string, conn *pgx.Conn) []byte {

	var jsonObj []byte
	err := conn.QueryRow(context.Background(), "SELECT json FROM orders WHERE id=$1", id).Scan(&jsonObj)
	if err != nil {
		log.Fatal(err)
	}
	return jsonObj

}

func GetInitialCache(conn *pgx.Conn) {

	cache := cache.NewCache()

	now := time.Now().Unix()
	period := now - 60*60*24
	//forCache := make([]string, 0, 10)

	query := `
	SELECT id, json
	FROM orders o
	WHERE o.time_of_creation > $1;`

	//query := `
	//SELECT id, json
	//FROM orders;
	//`

	//rows, err := conn.Query(context.Background(), query)
	rows, err := conn.Query(context.Background(), query, period)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//forCache := make([]string, 0, 10)

	for rows.Next() {

		var id *string
		var json *string

		err = rows.Scan(
			&id,
			&json,
		)

		if err != nil {
			log.Fatal(err)
		}
		cache.PutOrder(*id, *json)

		//forCache = append(forCache, *json)
		//log.Printf("forCache: %v", forCache)
	}

	//log.Printf(cache.GetOrder("b563feb7b2b84b6test"))

	log.Fatal("stop")

	return
}
