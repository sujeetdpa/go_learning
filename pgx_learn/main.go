package main

import (
	"github.com/sujeetdpa/go_learning/pgx_learn/controllers"
	"github.com/sujeetdpa/go_learning/pgx_learn/db"
)

func main() {
	db.PostgresSQLDatabaseConnection()
	controllers.GetAllAppUser()
}
