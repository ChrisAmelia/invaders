package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// resetStateOfAliens sets the aliens to nil.
// For testing purpose.
func resetStateOfAliens() {
	aliens = nil
}

// TestAlienDefaultAlien tests method DefaultAlien(int, int)
// generates a new alien if given positions (x,y) are not occupied,
// otherwise returns the existing alien.
func TestAlienDefaultAlien(t *testing.T) {
	resetStateOfAliens()

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

// TestAlienGetAlienAt tests method GetAlienAt(int, int)
// returns an alien located at given positions if it exists,
// otherwise tests than an error is returned.
func TestAlienGetAlienAt(t *testing.T) {
	resetStateOfAliens()

	x0, y0 := 0, 0
	x1, y1 := 1, 1

	// No alien has been generated on (0,0) yet,
	// expecting an error
	_, error := GetAlienAt(x0, y0)

	assert.NotNil(t, error)

	// Generating an alien at (0,0)
	// Expects an alien at (0,0) and no error (nil)
	DefaultAlien(x0, y0)

	alien0, error := GetAlienAt(x0, y0)
	assert.Nil(t, error)
	assert.Equal(t, x0, alien0.x)
	assert.Equal(t, y0, alien0.y)

	// Generating an alien at (1,1)
	// Expects an alien at (1,1) and no error (nil)
	DefaultAlien(x1, y1)

	alien1, error := GetAlienAt(x1, y1)

	assert.Nil(t, error)
	assert.Equal(t, x1, alien1.x)
	assert.Equal(t, y1, alien1.y)
}
