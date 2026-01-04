package services

import (
	"github.com/danyukod/joken-ia/internal/core/domain"
	"github.com/danyukod/joken-ia/internal/core/ports"
)

type gameService struct {
	aiProvider ports.AIProvider
}

func NewGameService(aiProvider ports.AIProvider) ports.GameService {
	return &gameService{
		aiProvider: aiProvider,
	}
}

func (s *gameService) Play(playerChoice domain.Choice) (domain.Round, error) {
	aiChoice := s.aiProvider.GetChoice()
	result := domain.GetResult(playerChoice, aiChoice)

	return domain.Round{
		PlayerChoice: playerChoice,
		AIChoice:     aiChoice,
		Result:       result,
	}, nil
}
