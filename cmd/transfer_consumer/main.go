package main

import (
	"context"
	"encoding/json"
	"log"
	"wallet/internal/postgres"
	"wallet/internal/rabbitmq"
	"wallet/internal/request"
)

func main() {
	db, err := postgres.NewConn()
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}
	defer db.Close(context.Background())

	amqp, err := rabbitmq.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer amqp.Close()

	q, err := rabbitmq.NewQueue(amqp,rabbitmq.Transfer)
	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()

	deliveries, err := q.Consume()
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}
	go func() {
		for d := range deliveries {
			r := &request.Transfer{}
			if err := json.Unmarshal(d.Body, &r); err != nil {
				log.Panicf("failed to bind delivery data %#v to request %s: %v", d.Body, rabbitmq.Transfer, err)
			}

			qry := `WITH sender AS (
				SELECT amount sAmount
				FROM wallets
				WHERE id = $1
			) UPDATE wallets SET amount = CASE
				WHEN id = $1 THEN amount - $3
				ELSE amount + $3
			END
			FROM sender
			WHERE sAmount >= $3 AND id IN ($1, $2)`
			commandTag, err := db.Exec(context.Background(), qry, r.Sender, r.Receiver, r.Amount)
			if err != nil {
				log.Panicf("transfer operation failed: %v", err)
			}
			if commandTag.RowsAffected() != 2 {
				log.Panic("transfer operation not executed properly!")
			}
		}
	}()
	<-forever
}
