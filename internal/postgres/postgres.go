package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// Connects to postgresql using os.Getenv params. Uses github.com/jackc/pgx/v4
func NewConn() (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DB"))
	return pgx.Connect(context.Background(), dsn)
}

func NewTestConn() (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_TEST_HOST"), os.Getenv("POSTGRES_TEST_PORT"), os.Getenv("POSTGRES_DB"))
	return pgx.Connect(context.Background(), dsn)
}
