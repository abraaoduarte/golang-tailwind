package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
    connection := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
    conn, err := pgx.Connect(context.Background(), connection)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
