package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"

	"dbut.dev/web-template/go/database"
	"dbut.dev/web-template/go/service"
)

func main() {
	dbConnString := "postgres://postgres:postgres@pg:5432/postgres?sslmode=disable"

	connConfig, err := pq.ParseURL(dbConnString)
	if err != nil {
		panic(err.Error())
	}
	conn, err := pq.NewConnector(connConfig)
	if err != nil {
		panic(err.Error())
	}
	db := sql.OpenDB(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err.Error())
	}

	svc := service.New(database.New(db))
	err = svc.Run(context.Background())
	if err != nil {
		panic(err.Error())
	}
}
