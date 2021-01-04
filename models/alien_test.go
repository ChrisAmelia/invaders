package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAlienDefaultAlien tests method DefaultAlien(int, int)
// generates a new alien if given positions (x,y) are not occupied,
// otherwise returns the existing alien.
func TestAlienDefaultAlien(t *testing.T) {
	x, y := 1, 1

	// Initial case: no aliens on board
	expected := 0
	actual := len(aliens)

	assert.Equal(t, expected, actual)

	// Generate one alien
	DefaultAlien(x, y)

	expected = 1
	actual = len(aliens)

	assert.Equal(t, expected, actual)

	// Generate one alien on the same position,
	// the number of aliens shouldn't increase
	DefaultAlien(x, y)

	expected = 1
	actual = len(aliens)

	assert.Equal(t, expected, actual)
}

// TestAlienMoveDown tests method moveDown()
// updates the alien's position.
func TestAlienMoveDown(t *testing.T) {
	alien := DefaultAlien(0, 0)

	alien.y = 7

	alien.MoveDown()

	expected := 6
	actual := alien.y

	assert.Equal(t, expected, actual)
}
