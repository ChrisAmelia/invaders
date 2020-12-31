package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAlienDefaultAlien tests method DefaultAlien(int, int)
// cannot generate two aliens on the same position.
func TestAlienDefaultAlien(t *testing.T) {
	x, y := 1, 1

	DefaultAlien(x, y)

	assert.Panics(t, func() { DefaultAlien(x, y) }, "Should panic when generating two aliens on the same positions.")
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
