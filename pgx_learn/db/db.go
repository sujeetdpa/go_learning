package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var PostresDB *pgx.Conn

func PostgresSQLDatabaseConnection() {
	// url := "postgres://<username>:<password>@ep-soft-unit-66642743-pooler.us-east-1.aws.neon.tech/vicdb"
	url := "postgres://vicadmin:XR0xLKW2kgYe@ep-soft-unit-66642743-pooler.us-east-1.aws.neon.tech/vicdb"
	conn, err := pgx.Connect(context.Background(), url)
	PostresDB = conn
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to PostgreSQL successfully :)")

	fmt.Println(greeting)
}
