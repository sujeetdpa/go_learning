package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5"
)

var PostresDB *pgx.Conn

// var DB *sql.DB

var DbHandle *goqu.Database

func PostgresSQLDatabaseConnection() {
	// url := "postgres://<username>:<password>@ep-soft-unit-66642743-pooler.us-east-1.aws.neon.tech/vicdb"
	// url := "postgres://vicadmin:XR0xLKW2kgYe@ep-soft-unit-66642743-pooler.us-east-1.aws.neon.tech/vicdb"
	url := "postgres://postgres:sujeet@localhost:5432/vicdb"
	conn, err := pgx.Connect(context.Background(), url)
	PostresDB = conn
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println("Unable to connect to PostgreSQL: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Connected to PostgreSQL successfully :)")
}

func PostgresSQLDatabaseConnectionUsingPGX() {
	db, err := sql.Open("pgx", "postgres://postgres:sujeet@localhost:5432/vicdb")
	if err != nil {
		log.Fatal(err)
	}
	// DB = db
	defer db.Close()

	DbHandle = goqu.New("pgx", db)

}
