package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/ariel-salgado/echo-sqlite-api/internal/server"
)

func main() {
	addr := os.Getenv("ADDR")

	s := server.NewServer(addr)

	if err := s.Start(); err != nil {
		log.Fatalf("Server could not be ran: %v", err)
	}
}