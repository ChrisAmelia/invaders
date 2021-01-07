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
