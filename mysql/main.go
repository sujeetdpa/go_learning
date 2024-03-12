package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	/*
		// Without using ORM
		db, err := sql.Open("mysql", "sujeet:sujeet@tcp(localhost:3306)/go")
		if err != nil {
			panic(err.Error())
		}
		pingerr := db.Ping()
		if pingerr != nil {
			panic(pingerr.Error())
		}
		fmt.Println("Successfull: Connected to mysql")

		//Creating a user table
		create, createerr := db.Query("create table user(id int,name varchar(20), primary key(id))")
		if createerr != nil {
			panic(createerr.Error())
		}
		defer create.Close()
	*/

	//Using ORM tool GORM
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}
