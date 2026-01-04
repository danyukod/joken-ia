package domain

import "testing"

func TestGetResult(t *testing.T) {
	tests := []struct {
		player Choice
		ai     Choice
		want   Result
	}{
		{Rock, Scissors, Win},
		{Rock, Paper, Loss},
		{Rock, Rock, Draw},
		{Paper, Rock, Win},
		{Paper, Scissors, Loss},
		{Paper, Paper, Draw},
		{Scissors, Paper, Win},
		{Scissors, Rock, Loss},
		{Scissors, Scissors, Draw},
	}

	for _, tt := range tests {
		t.Run(string(tt.player)+" vs "+string(tt.ai), func(t *testing.T) {
			if got := GetResult(tt.player, tt.ai); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
