package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAlienMoveDown tests method moveDown()
// updates the alien's position.
func TestAlienMoveDown(t *testing.T) {
	alien := DefaultAlien()

	alien.y = 7

	alien.MoveDown()

	expected := 6
	actual := alien.y

	assert.Equal(t, expected, actual)
}
