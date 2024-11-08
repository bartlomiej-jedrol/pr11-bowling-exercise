// main handles the Bowling Game Score CLI game.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MaxFrames int = 10
	MaxPins   int = 10
)

type Game struct {
	Frames []Frame
	Score  int
}

type Frame struct {
	Rolls      []int
	IsStrike   bool
	IsSpare    bool
	Score      int
	FinalScore int
}

// main handles the Bowling Game Score CLI game.
func main() {
	game := Game{}
	reader := bufio.NewReader(os.Stdin)

	for frameIndex := 0; frameIndex < MaxFrames; frameIndex++ {
		frame := Frame{}
		fmt.Printf("Enter rolls for frame %d (separated by space): ", frameIndex+1)
		input, _ := reader.ReadString('\n')
		rolls := strings.Fields(input)

		// Validate rolls
		isValidFrame := true
		for _, rollStr := range rolls {
			roll, err := strconv.Atoi(rollStr)
			if err != nil || roll < 0 || roll > MaxPins {
				fmt.Println("Invalid input. Please enter a number between 0 and 10.")
				frameIndex-- // Retry the current frame
				isValidFrame = false
				break
			}
			frame.Rolls = append(frame.Rolls, roll)
			frame.Score += roll
		}

		if !isValidFrame {
			continue
		}

		// Check if the total score of the frame exceeds 10 (except for the 10th frame)
		if frameIndex < MaxFrames-1 && frame.Score > MaxPins {
			fmt.Println("Invalid input. The total score for a frame cannot exceed 10.")
			frameIndex--
			continue
		}

		// Check for the number of rolls in frames 1-9
		if frameIndex < MaxFrames-1 && len(frame.Rolls) > 2 {
			fmt.Println("Invalid input. A frame can have a maximum of 2 rolls.")
			frameIndex--
			continue
		}

		// Check for the number of rolls in the 10th frame
		if frameIndex == MaxFrames-1 && len(frame.Rolls) > 3 {
			fmt.Println("Invalid input. The 10th frame can have a maximum of 3 rolls.")
			frameIndex--
			continue
		}

		if len(frame.Rolls) == 1 && frame.Score == MaxPins {
			frame.IsStrike = true
		} else if len(frame.Rolls) == 2 && frame.Score == MaxPins {
			frame.IsSpare = true
		}

		game.Frames = append(game.Frames, frame)

		calculateCurrentScore(&game)
		fmt.Printf("Current Score after frame %d: %d\n", frameIndex+1, game.Score)
	}

	fmt.Printf("\nFinal Score: %d\n", game.Score)
	fmt.Print("\nPress Enter to exit...")
	reader.ReadString('\n')
}

// calculateCurrentScore calcuates the current Bowling Game score.
func calculateCurrentScore(game *Game) {
	game.Score = 0
	runningTotal := 0
	for i := 0; i < len(game.Frames); i++ {
		frameScore := game.Frames[i].Score

		// Add strike bonus: next 2 rolls
		if i < MaxFrames-1 && game.Frames[i].IsStrike {
			frameScore += getNextTwoRollsScore(game, i)
		}

		// Add spare bonus: next 1 roll
		if i < MaxFrames-1 && game.Frames[i].IsSpare {
			frameScore += getNextRollScore(game, i)
		}

		runningTotal += frameScore
		game.Frames[i].FinalScore = runningTotal
		game.Score = runningTotal
	}

	// Display frame-by-frame scores
	fmt.Println("\nFrame-by-frame scores:")
	for i, frame := range game.Frames {
		fmt.Printf("Frame %d: %d", i+1, frame.FinalScore)
		if frame.IsStrike {
			fmt.Print(" (Strike!)")
		} else if frame.IsSpare {
			fmt.Print(" (Spare!)")
		}
		fmt.Println()
	}
}

// getNextTwoRollsScore returns the sum of the next two rolls after the given frame index.
func getNextTwoRollsScore(game *Game, currentIndex int) int {
	nextRolls := 0
	if currentIndex+1 < len(game.Frames) {
		nextRolls += game.Frames[currentIndex+1].Score
		if game.Frames[currentIndex+1].IsStrike && currentIndex+2 < len(game.Frames) {
			nextRolls += game.Frames[currentIndex+2].Rolls[0]
		}
	}
	return nextRolls
}

// getNextRollScore returns the score of the next roll after the given frame index.
func getNextRollScore(game *Game, currentIndex int) int {
	if currentIndex+1 < len(game.Frames) {
		return game.Frames[currentIndex+1].Rolls[0]
	}
	return 0
}
