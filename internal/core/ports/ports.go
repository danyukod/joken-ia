package ports

import "github.com/danyukod/joken-ia/internal/core/domain"

type GameService interface {
	Play(playerChoice domain.Choice) (domain.Round, error)
}

type AIProvider interface {
	GetChoice() domain.Choice
}
