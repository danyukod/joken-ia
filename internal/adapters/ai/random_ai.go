package ai

import (
	"math/rand"
	"time"

	"github.com/danyukod/joken-ia/internal/core/domain"
)

type randomAI struct{}

func NewRandomAI() *randomAI {
	rand.Seed(time.Now().UnixNano())
	return &randomAI{}
}

func (a *randomAI) GetChoice() domain.Choice {
	choices := []domain.Choice{domain.Rock, domain.Paper, domain.Scissors}
	return choices[rand.Intn(len(choices))]
}
