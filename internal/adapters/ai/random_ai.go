package ai

import (
	"joken-ia/internal/core/domain"
	"math/rand"
	"time"
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
