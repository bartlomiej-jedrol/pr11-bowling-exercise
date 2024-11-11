// main handles the Bowling Game Score CLI game.
package main

import (
	"testing"

	b "github.com/bartlomiej-jedrol/pr11-bowling-exercise/bowling"
)

// TestCalculateCurrentScore verifies the bowling game scoring logic through various scenarios
// including perfect games, spares, strikes, and open frames.
// For consistent testing, it uses predefined structs to inject predetermined framees wit roll values
// instead of user inputs. I covers negative scenarios to test validations.
func TestCalculateCurrentScore(t *testing.T) {
	tests := []struct {
		name          string
		frames        []b.Frame
		expectedScore int
	}{
		{
			name: "Perfect game",
			frames: []b.Frame{
				{Rolls: []int{10}, IsStrike: true, Score: 10},
				{Rolls: []int{10}, IsStrike: true, Score: 10},
				{Rolls: []int{10}, IsStrike: true, Score: 10},
				{Rolls: []int{10}, IsStrike: true, Score: 10},
				{Rolls: []int{10}, IsStrike: true, Score: 10},
				{Rolls: []int{10}, IsStrike: true, Score: 10},
				{Rolls: []int{10}, IsStrike: true, Score: 10},
				{Rolls: []int{10}, IsStrike: true, Score: 10},
				{Rolls: []int{10}, IsStrike: true, Score: 10},
				{Rolls: []int{10, 10, 10}, IsStrike: true, Score: 30},
			},
			expectedScore: 300,
		},
		{
			name: "All spares with 5s",
			frames: []b.Frame{
				{Rolls: []int{5, 5}, IsSpare: true, Score: 10},
				{Rolls: []int{5, 5}, IsSpare: true, Score: 10},
				{Rolls: []int{5, 5}, IsSpare: true, Score: 10},
				{Rolls: []int{5, 5}, IsSpare: true, Score: 10},
				{Rolls: []int{5, 5}, IsSpare: true, Score: 10},
				{Rolls: []int{5, 5}, IsSpare: true, Score: 10},
				{Rolls: []int{5, 5}, IsSpare: true, Score: 10},
				{Rolls: []int{5, 5}, IsSpare: true, Score: 10},
				{Rolls: []int{5, 5}, IsSpare: true, Score: 10},
				{Rolls: []int{5, 5, 5}, IsSpare: true, Score: 15},
			},
			expectedScore: 150,
		},
		{
			name: "Open frames",
			frames: []b.Frame{
				{Rolls: []int{3, 4}, Score: 7},
				{Rolls: []int{5, 4}, Score: 9},
				{Rolls: []int{3, 3}, Score: 6},
				{Rolls: []int{4, 5}, Score: 9},
				{Rolls: []int{3, 3}, Score: 6},
				{Rolls: []int{1, 5}, Score: 6},
				{Rolls: []int{3, 4}, Score: 7},
				{Rolls: []int{2, 1}, Score: 3},
				{Rolls: []int{4, 5}, Score: 9},
				{Rolls: []int{1, 2}, Score: 3},
			},
			expectedScore: 65,
		},
		{
			name: "All nines",
			frames: []b.Frame{
				{Rolls: []int{9, 0}, Score: 9},
				{Rolls: []int{9, 0}, Score: 9},
				{Rolls: []int{9, 0}, Score: 9},
				{Rolls: []int{9, 0}, Score: 9},
				{Rolls: []int{9, 0}, Score: 9},
				{Rolls: []int{9, 0}, Score: 9},
				{Rolls: []int{9, 0}, Score: 9},
				{Rolls: []int{9, 0}, Score: 9},
				{Rolls: []int{9, 0}, Score: 9},
				{Rolls: []int{9, 0}, Score: 9},
			},
			expectedScore: 90,
		},
		{
			name: "Last frame spare",
			frames: []b.Frame{
				{Rolls: []int{1, 2}, Score: 3},
				{Rolls: []int{3, 4}, Score: 7},
				{Rolls: []int{2, 3}, Score: 5},
				{Rolls: []int{4, 2}, Score: 6},
				{Rolls: []int{5, 1}, Score: 6},
				{Rolls: []int{2, 3}, Score: 5},
				{Rolls: []int{3, 4}, Score: 7},
				{Rolls: []int{2, 1}, Score: 3},
				{Rolls: []int{4, 2}, Score: 6},
				{Rolls: []int{7, 3, 5}, IsSpare: true, Score: 15},
			},
			expectedScore: 63,
		},
		{
			name: "Last frame strike",
			frames: []b.Frame{
				{Rolls: []int{1, 2}, Score: 3},
				{Rolls: []int{3, 4}, Score: 7},
				{Rolls: []int{2, 3}, Score: 5},
				{Rolls: []int{4, 2}, Score: 6},
				{Rolls: []int{5, 1}, Score: 6},
				{Rolls: []int{2, 3}, Score: 5},
				{Rolls: []int{3, 4}, Score: 7},
				{Rolls: []int{2, 1}, Score: 3},
				{Rolls: []int{4, 2}, Score: 6},
				{Rolls: []int{10, 3, 5}, IsStrike: true, Score: 18},
			},
			expectedScore: 66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := b.Game{Frames: tt.frames}
			b.CalculateFinalScore(&game)

			if game.Score != tt.expectedScore {
				t.Errorf("Expected score: %v, PlayGame() score: %v", tt.expectedScore, game.Score)
			}
		})
	}
}

// TestFrameValidation tests the validation logic for frames.
func TestFrameValidation(t *testing.T) {
	tests := []struct {
		name          string
		frame         b.Frame
		frameIndex    int
		expectedValid bool
	}{
		{
			name:          "Valid regular frame with strike",
			frame:         b.Frame{Rolls: []int{10}, Score: 10},
			frameIndex:    0,
			expectedValid: true,
		},
		{
			name:          "Valid regular frame with spare",
			frame:         b.Frame{Rolls: []int{5, 5}, Score: 10},
			frameIndex:    0,
			expectedValid: true,
		},
		{
			name:          "Valid tenth frame with strike",
			frame:         b.Frame{Rolls: []int{10, 10, 10}, Score: 30},
			frameIndex:    9,
			expectedValid: true,
		},
		{
			name:          "Valid tenth frame with spare",
			frame:         b.Frame{Rolls: []int{5, 5, 5}, Score: 15},
			frameIndex:    9,
			expectedValid: true,
		},
		{
			name:          "Invalid negative roll",
			frame:         b.Frame{Rolls: []int{-1}, Score: -1},
			frameIndex:    0,
			expectedValid: false,
		},
		{
			name:          "Invalid pin count",
			frame:         b.Frame{Rolls: []int{11}, Score: 11},
			frameIndex:    0,
			expectedValid: false,
		},
		{
			name:          "Invalid frame total",
			frame:         b.Frame{Rolls: []int{7, 8}, Score: 15},
			frameIndex:    0,
			expectedValid: false,
		},
		{
			name:          "Too many rolls in regular frame",
			frame:         b.Frame{Rolls: []int{1, 2, 3}, Score: 6},
			frameIndex:    0,
			expectedValid: false,
		},
		{
			name:          "Invalid tenth frame too many rolls",
			frame:         b.Frame{Rolls: []int{10, 10, 10, 10}, Score: 40},
			frameIndex:    9,
			expectedValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid, _ := validateFrame(tt.frame, tt.frameIndex)
			if isValid != tt.expectedValid {
				t.Errorf("Expected validation result %v, got %v", tt.expectedValid, isValid)
			}
		})
	}
}
