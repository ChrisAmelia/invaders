package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPlayerSetLife tests method SetLife
// updates the player's life.
func TestPlayerSetLife(t *testing.T) {
	player := DefaultPlayer()

	player.SetLife(14)

	expected := 14
	actual := player.life

	assert.Equal(t, expected, actual)
}

// TestPlayerSetX tests method SetX
// updates the player's abscissa.
func TestPlayerSetX(t *testing.T) {
	player := DefaultPlayer()

	player.SetX(14)

	expected := 14
	actual := player.x

	assert.Equal(t, expected, actual)
}

// TestPlayerSetUnicode tests method SetUnicode
// updates the player's unicode.
func TestPlayerSetUnicode(t *testing.T) {
	player := DefaultPlayer()

	player.SetUnicode("foobar")

	expected := "foobar"
	actual := player.unicode

	assert.Equal(t, expected, actual)
}

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

	// Player suddenly dies
	player.Die()

	assert.True(t, player.IsDead())
}

// TestPlayerChangeIconToSkull tests method ChangeIconToSkull()
// changes player's unicode when player is dead.
func TestPlayerChangeIconToSkull(t *testing.T) {
	player := DefaultPlayer()

	// Initial life: 3
	player.ChangeIconToSkull()

	expected := "Λ"
	actual := player.unicode

	assert.Equal(t, expected, actual)

	// Player got hit thrice, player is dead
	player.GetHit()
	player.GetHit()
	player.GetHit()
	player.ChangeIconToSkull()

	expected = "☠"
	actual = player.unicode

	assert.Equal(t, expected, actual)

}

// TestPlayerDie tests method Die()
// sets player's life to zero.
func TestPlayerDie(t *testing.T) {
	player := DefaultPlayer()

	// Initial life: 3
	expected := 3
	actual := player.life

	assert.Equal(t, expected, actual)

	// Player suddenly dies
	player.Die()

	expected = 0
	actual = player.life

	assert.Equal(t, expected, actual)
}
