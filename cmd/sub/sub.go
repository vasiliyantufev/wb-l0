package main

import (
	"flag"
	"github.com/gofrs/uuid"
	stan "github.com/nats-io/stan.go"
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

	// Subscribe to the ECHO channel as a queue.
	// Start with new messages as they come in; don't replay earlier messages.
	sub, err := sc.QueueSubscribe("ECHO", *queueGroup, func(msg *stan.Msg) {
		log.Printf("%10s | %s\n", msg.Subject, string(msg.Data))
	}, stan.StartWithLastReceived())
	if err != nil {
		log.Fatal(err)
	}

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
