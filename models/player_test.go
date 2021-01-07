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

// TestPlayerGetHit tests method GetHit()
// decreases player's life by one.
func TestPlayerGetHit(t *testing.T) {
	player := DefaultPlayer()

	// Initial life: 3
	expected := 3
	actual := player.life

	assert.Equal(t, expected, actual)

	// Player got hit, life decreases to 2
	player.GetHit()

	expected = 2
	actual = player.life

	assert.Equal(t, expected, actual)

	// Player got hit again, life decreases to 1
	player.GetHit()

	expected = 1
	actual = player.life

	assert.Equal(t, expected, actual)

	// Player got hit again, life decreases to 0
	player.GetHit()

	expected = 0
	actual = player.life

	assert.Equal(t, expected, actual)
}

// TestPlayerIsDead tests method IsDead()
// returns true if player's life reaches 0.
func TestPlayerIsDead(t *testing.T) {
	player := DefaultPlayer()

	// Initial life: 3
	assert.False(t, player.IsDead())

	// Player got hit, life decreases to 2
	player.GetHit()

	assert.False(t, player.IsDead())

	// Player got hit again, life decreases to 1
	player.GetHit()

	assert.False(t, player.IsDead())

	// Player got hit again, life decreases to 0
	player.GetHit()

	assert.True(t, player.IsDead())
}
