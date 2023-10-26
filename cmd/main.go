package main

import (
	db "athe-autho-go/database"
	"athe-autho-go/internal/logs"
	"athe-autho-go/internal/server"
	"athe-autho-go/migrations"
)

func main() {
	logger := logs.NewLogger()
	connection, dberr := db.Init()
	if dberr != nil {
		logger.Fatal("Could not start database", dberr)
	}

	migrations.Migrate(connection)
	defer connection.Close()

	r := server.NewServer()
	// Run the server
	r.Run(":8080")
}
