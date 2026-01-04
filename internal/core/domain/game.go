package domain

type Choice string

const (
	Rock     Choice = "rock"
	Paper    Choice = "paper"
	Scissors Choice = "scissors"
)

type Result string

const (
	Win  Result = "win"
	Loss Result = "loss"
	Draw Result = "draw"
)

type Round struct {
	PlayerChoice Choice
	AIChoice     Choice
	Result       Result
}

func GetResult(player, ai Choice) Result {
	if player == ai {
		return Draw
	}
	if (player == Rock && ai == Scissors) ||
		(player == Paper && ai == Rock) ||
		(player == Scissors && ai == Paper) {
		return Win
	}
	return Loss
}
