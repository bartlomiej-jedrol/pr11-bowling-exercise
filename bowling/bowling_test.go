// bowling handles the Bowling Game Score API game.
package bowling

import (
	"encoding/json"
	"testing"
)

// mockRollScore replaces the random roll generator with predetermined values.
func mockRollScore(rolls []int) func(n int) int {
	index := 0
	return func(n int) int {
		if index >= len(rolls) {
			panic("Not enough mock rolls provided")
		}
		roll := rolls[index]
		index++
		return roll
	}
}

// TestPlayGame verifies the bowling game scoring logic through various scenarios including
// perfect games, spares, strikes, and open frames. It uses mockRollScore to inject
// predetermined roll values instead of random ones for consistent testing. I does not cover
// negative scenarios as these are excluded by the bowling game scoring logic (random roll
// values are always between 1 an 10).
func TestPlayGame(t *testing.T) {
	tests := []struct {
		name          string
		rolls         []int
		expectedError bool
		expectedScore int
	}{
		{
			name:          "Perfect game",
			rolls:         []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			expectedScore: 300,
		},
		{
			name:          "All spares with 5s",
			rolls:         []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
			expectedScore: 150,
		},
		{
			name:          "Open frames",
			rolls:         []int{3, 4, 5, 4, 3, 3, 4, 5, 3, 3, 1, 5, 3, 4, 2, 1, 4, 5, 1, 2},
			expectedScore: 65,
		},
		{
			name:          "All nines",
			rolls:         []int{9, 0, 9, 0, 9, 0, 9, 0, 9, 0, 9, 0, 9, 0, 9, 0, 9, 0, 9, 0},
			expectedScore: 90,
		},
		{
			name:          "Last frame spare",
			rolls:         []int{1, 2, 3, 4, 2, 3, 4, 2, 5, 1, 2, 3, 3, 4, 2, 1, 4, 2, 7, 3, 5},
			expectedScore: 63,
		},
		{
			name:          "Last frame strike",
			rolls:         []int{1, 2, 3, 4, 2, 3, 4, 2, 5, 1, 2, 3, 3, 4, 2, 1, 4, 2, 10, 3, 5},
			expectedScore: 66,
		},
	}

	// Store the original rollScore function
	originalRollScore := rollScore

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace rollScore with mock
			rollScore = mockRollScore(tt.rolls)

			// Reset rollScore after the test
			defer func() {
				rollScore = originalRollScore
			}()

			result, _ := PlayGame()
			game := Game{}
			json.Unmarshal([]byte(result), &game)

			if game.Score != tt.expectedScore {
				t.Errorf("Expected score: %v, PlayGame() score: %v", tt.expectedScore, game.Score)
			}
		})
	}
}
