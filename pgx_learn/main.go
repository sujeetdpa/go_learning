package main

import (
	"github.com/sujeetdpa/go_learning/pgx_learn/db"
)

func main() {
	// db.PostgresSQLDatabaseConnection()
	// controllers.GetAllAppUser()
	db.PostgresSQLDatabaseConnectionUsingGoqu()
	// db.GoquDb.From("appuser").Executor().ScanStructs()
}
