// bowling handles the Bowling Game Score API game.
package bowling

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
)

const (
	MaxFrames int = 10
	MaxPins   int = 10
)

type Game struct {
	Frames []Frame `json:"frames"`
	Score  int     `json:"game_score"`
}

type Frame struct {
	Rolls    []int `json:"rolls"`
	IsStrike bool  `json:"is_strike"`
	IsSpare  bool  `json:"is_spare"`
	Score    int   `json:"frame_score"`
}

// rollScore holds a roll score function - returns random integer between 0 and 10.
var rollScore = func(n int) int {
	return rand.Intn(n + 1)
}

// getNextTwoRollsScore returns the sum of the next two rolls after the given frame index.
func GetNextTwoRollsScore(game *Game, frameIndex int) int {
	// Return 0 if we're at the last frame
	if frameIndex+1 >= len(game.Frames) {
		return 0
	}

	score := game.Frames[frameIndex+1].Rolls[0]

	// Get second roll either from next frame's second roll or the following frame's first roll
	if game.Frames[frameIndex+1].IsStrike && frameIndex+2 < len(game.Frames) {
		score += game.Frames[frameIndex+2].Rolls[0]
	} else if len(game.Frames[frameIndex+1].Rolls) > 1 {
		score += game.Frames[frameIndex+1].Rolls[1]
	}

	return score
}

// getNextRollScore returns the score of the next roll after the given frame index.
func GetNextRollScore(game *Game, frameIndex int) int {
	if frameIndex+1 >= len(game.Frames) {
		return 0
	}
	return game.Frames[frameIndex+1].Rolls[0]
}

// calculateFinalScore calcuates the final Bowling Game score.
func CalculateFinalScore(game *Game) {
	game.Score = 0
	for i := 0; i < len(game.Frames); i++ {
		frameScore := game.Frames[i].Score

		// Add strike bonus: next 2 rolls
		if i < MaxFrames-1 && game.Frames[i].IsStrike {
			frameScore += GetNextTwoRollsScore(game, i)
		}

		// Add spare bonus: next 1 roll
		if i < MaxFrames-1 && game.Frames[i].IsSpare {
			frameScore += GetNextRollScore(game, i)
		}

		game.Score += frameScore
	}
}

// marshalJSON marshals JSON of Game struct.
func marshalJSON(game Game) (string, error) {
	jsonData, err := json.Marshal(game)
	if err != nil {
		log.Printf("ERROR: PlayGame - %v", err)
		return "", err
	}
	return string(jsonData), nil
}

// PlayGame handles the Bowling Game Score API game.
func PlayGame() (string, error) {
	game := Game{}
	// globalRollIndex := 0

	for frameIndex := 0; frameIndex < MaxFrames; frameIndex++ {
		frame := Frame{}

		// First roll
		roll1 := rollScore(MaxPins)
		frame.Rolls = append(frame.Rolls, roll1)
		frame.Score = roll1

		// Strike handling
		if roll1 == MaxPins {
			frame.IsStrike = true

			// Special handling for 10th frame strike
			if frameIndex == MaxFrames-1 {
				// Two bonus rolls for 10th frame strike
				bonusRoll1 := rollScore(MaxPins)
				frame.Rolls = append(frame.Rolls, bonusRoll1)
				frame.Score += bonusRoll1

				remainingPins := MaxPins
				if bonusRoll1 < MaxPins {
					remainingPins = MaxPins - bonusRoll1
				}
				bonusRoll2 := rollScore(remainingPins)
				frame.Rolls = append(frame.Rolls, bonusRoll2)
				frame.Score += bonusRoll2
			}
			game.Frames = append(game.Frames, frame)
			continue
		}

		// Second roll if not a strike
		remainingPins := MaxPins - roll1
		roll2 := rollScore(remainingPins)
		frame.Rolls = append(frame.Rolls, roll2)
		frame.Score += roll2

		// Spare handling
		if frame.Score == MaxPins {
			frame.IsSpare = true
			// Special handling for 10th frame spare
			if frameIndex == MaxFrames-1 {
				bonusRoll1 := rollScore(MaxPins)
				frame.Rolls = append(frame.Rolls, bonusRoll1)
				frame.Score += bonusRoll1
			}
		}

		game.Frames = append(game.Frames, frame)
	}
	for _, frame := range game.Frames {
		fmt.Printf("frame: %v\n", frame)
	}
	fmt.Printf("game: %v", game)
	CalculateFinalScore(&game)

	jsonString, err := marshalJSON(game)
	return jsonString, err
}
