package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	stan "github.com/nats-io/stan.go"
	"github.com/vasiliyantufev/wb-l0/internal/app"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var (
		clusterID  = flag.String("cluster_id", "test-cluster", "Cluster ID")
		clientID   = flag.String("client_id", "", "Client ID")
		queueGroup = flag.String("queue-group", "", "Queue group ID")
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

	conn, err := pgx.Connect(context.Background(), "postgres://root:password@localhost:5532/wb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Subscribe to the ECHO channel as a queue.
	// Start with new messages as they come in; don't replay earlier messages.
	sub, err := sc.QueueSubscribe("ECHO", *queueGroup, func(msg *stan.Msg) {
		//_, _ = app.ParseMessages(msg.Data)

		message, err := app.ParseMessages(msg.Data)
		if err != nil {
			log.Println("no data")
		} else {

			app.GetInitialCache(conn)
			//log.Println(message)

			//cache, err := repo.GetInitialCache(ctx)
			//cache.PutOrder(message.OrderUid, string(msg.Data))

			err = app.InsertOrder(message, conn)
		}

	}, stan.StartWithLastReceived())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("start")
	// Wait for Ctrl+C
	doneCh := make(chan bool)
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt)
		<-sigCh
		sub.Unsubscribe()
		doneCh <- true
	}()
	<-doneCh
}
