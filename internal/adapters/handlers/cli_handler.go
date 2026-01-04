package handlers

import (
	"fmt"
	"joken-ia/internal/core/domain"
	"joken-ia/internal/core/ports"
	"strings"
)

type CLIHandler struct {
	gameService ports.GameService
}

func NewCLIHandler(gs ports.GameService) *CLIHandler {
	return &CLIHandler{gameService: gs}
}

func (h *CLIHandler) Start() {
	fmt.Println("Welcome to Jokenpo with AI!")
	fmt.Println("Options: rock, paper, scissors (or 'exit' to quit)")

	for {
		fmt.Print("\nYour move: ")
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		if input == "exit" {
			break
		}

		choice := domain.Choice(input)
		if choice != domain.Rock && choice != domain.Paper && choice != domain.Scissors {
			fmt.Println("Invalid choice. Try again.")
			continue
		}

		round, err := h.gameService.Play(choice)
		if err != nil {
			fmt.Printf("Error playing: %v\n", err)
			continue
		}

		fmt.Printf("AI chose: %s\n", round.AIChoice)
		fmt.Printf("Result: %s\n", round.Result)
	}
}
