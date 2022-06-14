package main

import (
	"context"
	"wallet/internal/postgres"
	"wallet/internal/rabbitmq"
)

func main() {
	s := newServer()
	db, err := postgres.NewConn()
	if err != nil {
		s.echo.Logger.Fatal(err)
	}
	defer db.Close(context.Background())

	amqp, err := rabbitmq.NewConn()
	if err != nil {
		s.echo.Logger.Fatal(err)
	}
	defer amqp.Close()

	s.db = db
	s.amqp = amqp

	s.echo.Logger.Fatal(s.echo.Start(":8080"))
}
