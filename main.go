package main

import (
	"joken-ia/internal/adapters/ai"
	"joken-ia/internal/adapters/handlers"
	"joken-ia/internal/core/services"
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
