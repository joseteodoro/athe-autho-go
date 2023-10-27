package main

import (
	"athe-autho-go/internal/configs"
	"athe-autho-go/internal/logs"
	"athe-autho-go/internal/migrations"
	"athe-autho-go/internal/server"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	logger := logs.NewLogger()

	migrations.Migrate()

	r := server.NewServer()
	// Run the server
	port := configs.NewApplicationConfig().Port
	logger.Info(strings.Join([]string{"Starting application at port:", port}, ""))
	r.Run(strings.Join([]string{":", port}, ""))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
}
