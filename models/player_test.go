package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPlayerMoveRight tests method moveRight()
// updates the player's position.
func TestPlayerMoveRight(t *testing.T) {
	player := DefaultPlayer()

	player.MoveRight()

	expected := 1
	actual := player.x

	assert.Equal(t, expected, actual)
}

// TestPlayerMoveLeft tests method moveLeft()
// updates the player's position.
func TestPlayerMoveLeft(t *testing.T) {
	player := DefaultPlayer()

	player.MoveLeft()

	expected := -1
	actual := player.x

	assert.Equal(t, expected, actual)
}
