// main handles the Bowling Game Score CLI game.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	b "github.com/bartlomiej-jedrol/pr11-bowling-exercise/bowling"
)

// validateFrame checks if a frame is valid according to bowling rules.
func validateFrame(frame b.Frame, frameIndex int) (bool, string) {
	// Check if the roll is negative.
	for _, roll := range frame.Rolls {
		if roll < 0 {
			return false, "Invalid input. A roll cannot be a negative value."
		}
	}

	// Check if the total score of the frame exceeds 10 (except for the 10th frame).
	if frameIndex < b.MaxFrames-1 && frame.Score > b.MaxPins {
		return false, "Invalid input. The total score for a frame cannot exceed 10."
	}

	// Check for the number of rolls in frames 1-9.
	if frameIndex < b.MaxFrames-1 && len(frame.Rolls) > 2 {
		return false, "Invalid input. A frame can have a maximum of 2 rolls."
	}

	// Check for the number of rolls in the 10th frame.
	if frameIndex == b.MaxFrames-1 && len(frame.Rolls) > 3 {
		return false, "Invalid input. The 10th frame can have a maximum of 3 rolls."
	}

	return true, ""
}

// main handles the Bowling Game Score CLI game.
func main() {
	game := b.Game{}
	reader := bufio.NewReader(os.Stdin)

	for frameIndex := 0; frameIndex < b.MaxFrames; frameIndex++ {
		frame := b.Frame{}
		fmt.Printf("Enter rolls for frame %d (separated by space): ", frameIndex+1)
		input, _ := reader.ReadString('\n')
		rolls := strings.Fields(input)

		// Validate rolls
		isValidFrame := true
		for _, rollStr := range rolls {
			roll, err := strconv.Atoi(rollStr)
			if err != nil || roll < 0 || roll > b.MaxPins {
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

		if valid, message := validateFrame(frame, frameIndex); !valid {
			fmt.Println(message)
			frameIndex--
			continue
		}

		if len(frame.Rolls) == 1 && frame.Score == b.MaxPins {
			frame.IsStrike = true
		} else if len(frame.Rolls) == 2 && frame.Score == b.MaxPins {
			frame.IsSpare = true
		}

		game.Frames = append(game.Frames, frame)

		b.CalculateFinalScore(&game)
		displayFrameScores(&game)
		fmt.Printf("Current Score after frame %d: %d\n", frameIndex+1, game.Score)
	}

	fmt.Printf("\nFinal Score: %d\n", game.Score)
	fmt.Print("\nPress Enter to exit...")
	reader.ReadString('\n')
}

// displayFrameScores displays frame-by-frame scores.
func displayFrameScores(g *b.Game) {
	fmt.Println("\nFrame-by-frame scores:")
	for i, frame := range g.Frames {
		fmt.Printf("Frame %d: %d", i+1, frame.Score)
		if frame.IsStrike {
			fmt.Print(" (Strike!)")
		} else if frame.IsSpare {
			fmt.Print(" (Spare!)")
		}
		fmt.Println()
	}
}
