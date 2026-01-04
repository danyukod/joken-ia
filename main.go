package main

import (
	"github.com/danyukod/joken-ia/internal/adapters/ai"
	"github.com/danyukod/joken-ia/internal/adapters/handlers"
	"github.com/danyukod/joken-ia/internal/core/services"
)

func main() {
	// Initialize Driven Adapter (AI)
	aiProvider := ai.NewRandomAI()

	// Initialize Core Service
	gameService := services.NewGameService(aiProvider)

	// Initialize Driver Adapter (CLI)
	cliHandler := handlers.NewCLIHandler(gameService)

	// Start application
	cliHandler.Start()
}
