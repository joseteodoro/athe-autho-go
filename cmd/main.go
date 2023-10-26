package main

import (
	"athe-autho-go/configs"
	db "athe-autho-go/database"
	"athe-autho-go/internal/logs"
	"athe-autho-go/internal/server"
	"athe-autho-go/migrations"
	"strings"
)

func main() {
	logger := logs.NewLogger()
	connection := db.Init()

	migrations.Migrate(connection)
	defer connection.Close()

	r := server.NewServer()
	// Run the server
	port := configs.NewApplicationConfig().Port
	logger.Info(strings.Join([]string{"Starting application at port:", port}, ""))
	r.Run(strings.Join([]string{":", port}, ""))
}
