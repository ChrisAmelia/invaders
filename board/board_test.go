package board

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNegativeDimensionsBoardString tests the string representation
// of the board with negative dimensions.
func TestNegativeDimensionsBoardString(t *testing.T) {
	board := NewBoard(-1, -1)

	assert.Panics(t, func() { fmt.Println(board.String()) }, "The code did not panic.")
}

// Test0x0BoardToString tests the string representation
// of the board with dimensions 0x0.
func Test0x0BoardString(t *testing.T) {
	ResetStateOfAliens()

	board := NewBoard(0, 0)

	expected := ""
	actual := board.String()

	assert.Equal(t, expected, actual, "Expecting an empty string for a board with no dimensions.")
}

// Test1x1BoardToString tests the string representation
// of the board with dimensions 1x1.
func Test1x1BoardString(t *testing.T) {
	ResetStateOfAliens()

	board := NewBoard(1, 1)

	expected := ""
	expected += "+-+" + "\n"
	expected += "| |" + "\n"
	expected += "*Λ*"

	actual := board.String()

	message := "Expecting a square, representation should be:\n"
	message += expected

	assert.Equal(t, expected, actual, message)
}

// Test4x4BoardString tests the string representation
// of the board with dimensions 4x4.
func Test4x4BoardString(t *testing.T) {
	ResetStateOfAliens()

	board := NewBoard(4, 4)

	expected := ""
	expected += "+----+" + "\n"
	expected += "|    |" + "\n"
	expected += "|    |" + "\n"
	expected += "|    |" + "\n"
	expected += "|    |" + "\n"
	expected += "*Λ   *"

	actual := board.String()

	message := "Expecting a square, representation should be:\n"
	message += expected

	assert.Equal(t, expected, actual, message)
}

// Test4x4BoardStringWithPlayerBottomRight tests the string representation
// of the board with dimensions 4x4 and a player present
// on the bottom right.
func Test4x4BoardStringWithPlayerBottomRight(t *testing.T) {
	ResetStateOfAliens()

	board := NewBoard(4, 4)

	board.player.SetX(board.width - 1)

	expected := ""
	expected += "+----+" + "\n"
	expected += "|    |" + "\n"
	expected += "|    |" + "\n"
	expected += "|    |" + "\n"
	expected += "|    |" + "\n"
	expected += "*   Λ*"

	actual := board.String()

	message := "Expecting a square, representation should be :\n"
	message += expected

	assert.Equal(t, expected, actual, message)
}


// Test5x5BoardStringWithPlayerBottomRight tests the string representation
// of the board with dimensions 5x5 and a player present
// on the middle.
func Test5x5BoardStringWithPlayerBottomRight(t *testing.T) {
	ResetStateOfAliens()

	board := NewBoard(5, 5)

	board.player.SetX(2)

	expected := ""
	expected += "+-----+" + "\n"
	expected += "|     |" + "\n"
	expected += "|     |" + "\n"
	expected += "|     |" + "\n"
	expected += "|     |" + "\n"
	expected += "|     |" + "\n"
	expected += "*  Λ  *"

	actual := board.String()

	message := "Expecting a square, representation should be :\n"
	message += expected

	assert.Equal(t, expected, actual, message)
}

// Test5x5BoardStringWithAliens tests the string representation
// of the board with dimensions 5x5 and aliens present
// over the board.
func Test5x5BoardStringWithAliens(t *testing.T) {
	ResetStateOfAliens()

	board := NewBoard(5, 5)

	GenerateAlienAt(0, 0)
	GenerateAlienAt(0, 4)
	GenerateAlienAt(4, 0)
	GenerateAlienAt(4, 4)

	expected := ""
	expected += "+-----+" + "\n"
	expected += "|V   V|" + "\n"
	expected += "|     |" + "\n"
	expected += "|     |" + "\n"
	expected += "|     |" + "\n"
	expected += "|V   V|" + "\n"
	expected += "*Λ    *"

	actual := board.String()

	message := "Expecting a square, representation should be :\n"
	message += expected

	assert.Equal(t, expected, actual, message)

}
