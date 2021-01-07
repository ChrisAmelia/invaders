package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPlayerMoveRight tests method MoveRight()
// updates the player's position.
func TestPlayerMoveRight(t *testing.T) {
	player := DefaultPlayer()

	player.MoveRight()

	expected := 1
	actual := player.x

	assert.Equal(t, expected, actual)
}

// TestPlayerMoveLeft tests method MoveLeft()
// updates the player's position.
func TestPlayerMoveLeft(t *testing.T) {
	player := DefaultPlayer()

	player.MoveLeft()

	expected := -1
	actual := player.x

	assert.Equal(t, expected, actual)
}

// TestPlayerGetHurt tests method GetHurts()
// decreases player's life by one.
func TestPlayerGetHurt(t *testing.T) {
	player := DefaultPlayer()

	// Initial life: 3
	expected := 3
	actual := player.life

	assert.Equal(t, expected, actual)

	// Player got hit, life decreases to 2
	player.GetHurt()

	expected = 2
	actual = player.life

	assert.Equal(t, expected, actual)

	// Player got hit again, life decreases to 1
	player.GetHurt()

	expected = 1
	actual = player.life

	assert.Equal(t, expected, actual)

	// Player got hit again, life decreases to 0
	player.GetHurt()

	expected = 0
	actual = player.life

	assert.Equal(t, expected, actual)
}
