// main scores Bowling Game.
package main

import (
	"fmt"
	"math/rand"
)

const (
	MaxFrames int = 10
	MaxPins   int = 10
)

type Game struct {
	ID     int
	Frames []Frame
	Score  int
}

type Frame struct {
	ID       int
	Rolls    []Roll
	IsStrike bool
	IsSpare  bool
	Score    int
}

type Roll struct {
	ID    int
	Score int
}

// AddGame adds a Bowling Game.
func AddGame(gameIndex int) Game {
	return Game{
		ID: gameIndex + 1,
	}
}

// AddFrame adds a frame.
func AddFrame(frameIndex int) Frame {
	return Frame{
		ID:    frameIndex + 1,
		Score: 0,
	}
}

// AddRoll adds a roll.
func AddRoll(globalRollIndex, pins int) Roll {
	return Roll{
		ID:    globalRollIndex + 1,
		Score: RollScore(pins),
	}
}

// RollScore returns roll score - random integer between 0 and 10.
func RollScore(n int) int {
	return rand.Intn(n + 1)
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
		fmt.Printf("Frame %d: %d\n", i+1, frameScore)
	}
	fmt.Printf("Final Score: %d\n", game.Score)
}

// main handles Bowling Game logic.
func main() {
	game := AddGame(0)
	globalRollIndex := 0

	for frameIndex := 0; frameIndex < MaxFrames; frameIndex++ {
		frame := AddFrame(frameIndex)
		game.Frames = append(game.Frames, frame)

		// First roll
		roll := AddRoll(globalRollIndex, MaxPins)
		globalRollIndex++
		frame.Rolls = append(frame.Rolls, roll)
		frame.Score = roll.Score

		// Strike handling
		if frame.Score == MaxPins {
			frame.IsStrike = true
			if frameIndex < MaxFrames-1 {
				game.Frames[frameIndex] = frame
				continue
			}
			// For 10th frame strike, skip the normal second roll
			// and go straight to bonus rolls
			roll = AddRoll(globalRollIndex, MaxPins)
			globalRollIndex++
			frame.Rolls = append(frame.Rolls, roll)
			frame.Score += roll.Score

			// Second bonus roll
			remainingPins := MaxPins
			if roll.Score < MaxPins {
				remainingPins = MaxPins - roll.Score
			}
			roll = AddRoll(globalRollIndex, remainingPins)
			globalRollIndex++
			frame.Rolls = append(frame.Rolls, roll)
			frame.Score += roll.Score

			game.Frames[frameIndex] = frame
			continue
		}

		// Second roll (only reaches here if not a strike)
		remainingPins := MaxPins - frame.Score
		roll = AddRoll(globalRollIndex, remainingPins)
		globalRollIndex++
		frame.Rolls = append(frame.Rolls, roll)
		frame.Score += roll.Score

		// Spare handling
		if frame.Score == MaxPins {
			frame.IsSpare = true
			if frameIndex == MaxFrames-1 {
				// One bonus roll for spare in 10th frame
				roll = AddRoll(globalRollIndex, MaxPins)
				globalRollIndex++
				frame.Rolls = append(frame.Rolls, roll)
				frame.Score += roll.Score
			}
		}

		game.Frames[frameIndex] = frame
	}
	for _, frame := range game.Frames {
		fmt.Printf("frame: %v\n", frame)
	}
	calculateFinalScore(&game)
}
