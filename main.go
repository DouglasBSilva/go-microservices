package main

import (
	"log"

	"github.com/DouglasBSilva/go-microservices/internal/database"
	"github.com/DouglasBSilva/go-microservices/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()

	if err != nil {
		log.Fatalf("Failed to initialize Database Client: %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
