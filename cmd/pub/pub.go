package main

import (
	"flag"
	"github.com/gofrs/uuid"
	stan "github.com/nats-io/stan.go"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var (
		clusterID = flag.String("cluster_id", "test-cluster", "Cluster ID")
		clientID  = flag.String("client_id", "", "Client ID")
	)
	flag.Parse()

	if *clientID == "" {
		*clientID = uuid.Must(uuid.NewV4()).String()
	}

	// Connect to NATS Streaming Server cluster
	sc, err := stan.Connect(*clusterID, *clientID,
		stan.NatsURL("nats://localhost:4222"),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Printf("Connection lost: %v", reason)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// Publish some messages
	test1 := `{"order_uid": "test4", 
	"track_number": "WBILMTESTTRACK", 
	"entry": "WBIL",
	"delivery": {
	  "name": "Test Testov",
	  "phone": "+9720000000",
	  "zip": "2639809",
	  "city": "Kiryat Mozkin",
	  "address": "Ploshad Mira 15",
	  "region": "Kraiot",
	  "email": "test@gmail.com"
	},
	"payment": {
	  "transaction": "b563feb7b2b84b6test",
	  "request_id": "",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 1817,
	  "payment_dt": 1637907727,
	  "bank": "alpha",
	  "delivery_cost": 1500,
	  "goods_total": 317,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 9934930,
		"track_number": "WBILMTESTTRACK",
		"price": 453,
		"rid": "ab4219087a764ae0btest",
		"name": "Mascaras",
		"sale": 30,
		"size": "0",
		"total_price": 317,
		"nm_id": 2389212,
		"brand": "Vivienne Sabo",
		"status": 202
	  }
	],
	"locale": "en",
	"internal_signature": "",
	"customer_id": "test",
	"delivery_service": "meest",
	"shardkey": "9",
	"sm_id": 99,
	"date_created": "2021-11-26T06:22:19Z",
	"oof_shard": "1"
  }`

	test2 := `{"order_uid": "test5", 
	"track_number": "WBILMTESTTRACK", 
	"entry": "WBIL",
	"delivery": {
	  "name": "Test Testov",
	  "phone": "+9720000000",
	  "zip": "2639809",
	  "city": "Kiryat Mozkin",
	  "address": "Ploshad Mira 15",
	  "region": "Kraiot",
	  "email": "test@gmail.com"
	},
	"payment": {
	  "transaction": "b563feb7b2b84b6test",
	  "request_id": "",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 1817,
	  "payment_dt": 1637907727,
	  "bank": "alpha",
	  "delivery_cost": 1500,
	  "goods_total": 317,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 9934930,
		"track_number": "WBILMTESTTRACK",
		"price": 453,
		"rid": "ab4219087a764ae0btest",
		"name": "Mascaras",
		"sale": 30,
		"size": "0",
		"total_price": 317,
		"nm_id": 2389212,
		"brand": "Vivienne Sabo",
		"status": 202
	  }
	],
	"locale": "en",
	"internal_signature": "",
	"customer_id": "test",
	"delivery_service": "meest",
	"shardkey": "9",
	"sm_id": 99,
	"date_created": "2021-11-26T06:22:19Z",
	"oof_shard": "1"
  }`

	test3 := `{"order_uid": "b563feb7b2b84b6test"}`

	test4 := `{"b133f2wb2qwf4b6test"}`

	test5 := `264gfdg43g4d4t34`

	log.Println("start writing")
	err = sc.Publish("ECHO", []byte(test1))
	if err != nil {
		log.Fatal(err)
	}
	err = sc.Publish("ECHO", []byte(test2))
	if err != nil {
		log.Fatal(err)
	}
	err = sc.Publish("ECHO", []byte(test3))
	if err != nil {
		log.Fatal(err)
	}
	err = sc.Publish("ECHO", []byte(test4))
	if err != nil {
		log.Fatal(err)
	}
	err = sc.Publish("ECHO", []byte(test5))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)

	//// Publish some messages, synchronously
	//var i int64
	//for {
	//	i++
	//	now := time.Now().Format(time.RFC3339)
	//	payload := fmt.Sprintf("%08d %s", i, now)
	//	err := sc.Publish("ECHO", []byte(payload))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	// Sleep for a random time of up to 1s
	//	time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
	//}
}
