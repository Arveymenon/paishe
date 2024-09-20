package main

import (
	"log"

	"github.com/Arveymenon/paishe/internal/config"
	"github.com/Arveymenon/paishe/internal/transport/http"
)

func main() {
	cfg := config.LoadConfig()

	config.SetUpDataBase()

	router := http.SetupRouter()
	log.Println("Starting server on port", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
