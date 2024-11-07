package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

const (
	MaxFrames int = 10
	MaxPins   int = 10
)

type Game struct {
	ID     int     `json:"id"`
	Frames []Frame `json:"frames"`
	Score  int     `json:"score"`
}

type Frame struct {
	ID       int    `json:"id"`
	Rolls    []Roll `json:"rolls"`
	IsStrike bool   `json:"is_strike"`
	IsSpare  bool   `json:"is_spare"`
	Score    int    `json:"score"`
}

type Roll struct {
	ID    int `json:"id"`
	Score int `json:"score"`
}

// newGame creates a Bowling Game.
func newGame(gameIndex int) Game {
	return Game{
		ID: gameIndex + 1,
	}
}

// newFrame creates a frame.
func newFrame(frameIndex int) Frame {
	return Frame{
		ID:    frameIndex + 1,
		Score: 0,
	}
}

// newRoll creates a roll.
func newRoll(globalRollIndex *int, pins int) Roll {
	*globalRollIndex++
	return Roll{
		ID:    *globalRollIndex,
		Score: rollScore(pins),
	}
}

// rollScore returns roll score - random integer between 0 and 10.
func rollScore(n int) int {
	return rand.Intn(n + 1)
}

// addRollToFrame adds roll to frame.
func addRollToFrame(globalRollIndex *int, game *Game, frameIndex, pins int) Frame {
	roll := newRoll(globalRollIndex, pins)

	frame := game.Frames[frameIndex]
	frame.Rolls = append(frame.Rolls, roll)
	frame.Score += roll.Score
	game.Frames[frameIndex] = frame
	return frame
}

// calculateFinalScore calculates Bowling Game final score
func calculateFinalScore(game *Game) {
	game.Score = 0
	for i := 0; i < len(game.Frames); i++ {
		frameScore := game.Frames[i].Score

		// Add strike bonus: next 2 rolls
		if i < MaxFrames-1 && game.Frames[i].IsStrike && i+1 < len(game.Frames) {
			frameScore += game.Frames[i+1].Rolls[0].Score // Add next roll

			// Add second bonus roll
			// (either from next frame's second roll or the following frame's first roll)
			if game.Frames[i+1].IsStrike && i+2 < len(game.Frames) {
				frameScore += game.Frames[i+2].Rolls[0].Score
			} else if len(game.Frames[i+1].Rolls) > 1 {
				frameScore += game.Frames[i+1].Rolls[1].Score
			}
		}

		// Add spare bonus: next 1 roll
		if i < MaxFrames-1 && game.Frames[i].IsSpare {
			if i+1 < len(game.Frames) {
				frameScore += game.Frames[i+1].Rolls[0].Score
			}
		}

		game.Score += frameScore
		// fmt.Printf("Frame %d: %d\n", i+1, frameScore)
	}
	// fmt.Printf("Final Score: %d\n", game.Score)
}

func main() {
	game := newGame(0)
	globalRollIndex := 0

	for frameIndex := 0; frameIndex < MaxFrames; frameIndex++ {
		frame := newFrame(frameIndex)
		game.Frames = append(game.Frames, frame)

		// First roll
		frame = addRollToFrame(&globalRollIndex, &game, frameIndex, MaxPins)
		// Strike handling
		if frame.Score == MaxPins {
			frame.IsStrike = true
			game.Frames[frameIndex] = frame
			if frameIndex < MaxFrames-1 {
				game.Frames[frameIndex] = frame
				continue
			}
			// For 10th frame strike, skip the normal second roll and go straight to bonus rolls
			frame = addRollToFrame(&globalRollIndex, &game, frameIndex, MaxPins)

			// Second bonus roll
			remainingPins := MaxPins
			// Calculate remaining pins based on previous roll score
			if frame.Rolls[len(frame.Rolls)-1].Score < MaxPins {
				remainingPins = MaxPins - frame.Rolls[len(frame.Rolls)-1].Score
			}
			frame = addRollToFrame(&globalRollIndex, &game, frameIndex, remainingPins)
			continue
		}

		// Second roll (only reaches here if not a strike)
		remainingPins := MaxPins - frame.Score
		frame = addRollToFrame(&globalRollIndex, &game, frameIndex, remainingPins)

		// Spare handling
		if frame.Score == MaxPins {
			frame.IsSpare = true
			game.Frames[frameIndex] = frame
			if frameIndex == MaxFrames-1 {
				// One bonus roll for spare in 10th frame
				frame = addRollToFrame(&globalRollIndex, &game, frameIndex, remainingPins)
			}
		}
	}
	// for _, frame := range game.Frames {
	// 	fmt.Printf("frame: %v\n", frame)
	// }
	calculateFinalScore(&game)
	// fmt.Printf("game: %v", game)
	gameString, err := json.Marshal(game)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Printf("game string: %v", string(gameString))
}
