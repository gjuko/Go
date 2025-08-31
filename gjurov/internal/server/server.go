package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"gjurov/internal/database"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // fallback if .env is missing or PORT is unset
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid PORT value: %v", err)
	}

	s := &Server{
		port: port,
		db:   database.New(),
	}

	log.Printf("Starting server on port %d", s.port)

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
