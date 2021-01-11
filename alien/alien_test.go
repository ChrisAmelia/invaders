package alien

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAlienSetLife tests method SetLife
// update the alien's life.
func TestAlienSetLife(t *testing.T) {
	ResetStateOfAliens()

	alien := DefaultAlien(0, 0)

	alien.SetLife(42)

	expected := 42
	actual := alien.life

	assert.Equal(t, expected, actual)
}

// TestAlienSetX tests method SetX
// update the alien's abscissa.
func TestAlienSetX(t *testing.T) {
	ResetStateOfAliens()

	alien := DefaultAlien(0, 0)

	alien.SetX(42)

	expected := 42
	actual := alien.x

	assert.Equal(t, expected, actual)
}

// TestAlienSetY tests method SetY
// update the alien's ordinate.
func TestAlienSetY(t *testing.T) {
	ResetStateOfAliens()

	alien := DefaultAlien(0, 0)

	alien.SetY(42)

	expected := 42
	actual := alien.y

	assert.Equal(t, expected, actual)
}

// TestAlienSetUnicode tests method SetUnicode
// update the alien's unicode.
func TestAlienSetUnicode(t *testing.T) {
	ResetStateOfAliens()

	alien := DefaultAlien(0, 0)

	alien.SetUnicode("foobar")

	expected := "foobar"
	actual := alien.unicode

	assert.Equal(t, expected, actual)
}
// TestAlienDefaultAlien tests method DefaultAlien(int, int)
// generates a new alien if given positions (x,y) are not occupied,
// otherwise returns the existing alien.
func TestAlienDefaultAlien(t *testing.T) {
	ResetStateOfAliens()

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
	ResetStateOfAliens()

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

// TestAlienGetHit tests method GetHit()
// decreases alien's life by one.
func TestAlienGetHit(t *testing.T) {
	ResetStateOfAliens()

	// Initial state
	alien := DefaultAlien(0, 0)

	expected := 1
	actual := alien.life

	assert.Equal(t, expected, actual)

	// Alien got hit, life decreases to 0
	alien.GetHit()

	expected = 0
	actual = alien.life

	assert.Equal(t, expected, actual)
}

// TestAlienChangeIconToEmpty tests method ChangeIconToEmpty()
// changes alien's unicode when alien is dead.
func TestAlienChangeIconToEmpty(t *testing.T) {
	ResetStateOfAliens()

	alien := DefaultAlien(0, 0)

	// Initial life: 1
	alien.ChangeIconToEmpty()

	expected := "V"
	actual := alien.unicode

	assert.Equal(t, expected, actual)

	// Alien got hit, alien is dead
	alien.GetHit()
	alien.ChangeIconToEmpty()

	expected = ""
	actual = alien.unicode

	assert.Equal(t, expected, actual)
}

// TestAlienIsDead tests method IsDead()
// returns true if alien's life has reached zero.
func TestAlienIsDead(t *testing.T) {
	ResetStateOfAliens()

	alien := DefaultAlien(0, 0)

	// Initial life: 1
	assert.False(t, alien.IsDead())

	// Alien got hit, life decreases to 0
	alien.GetHit()

	assert.True(t, alien.IsDead())
}

// TestAlienDie tests method Die()
// sets alien's life to zero.
func TestAlienDie(t *testing.T) {
	ResetStateOfAliens()

	alien := DefaultAlien(0, 0)

	// Initial life: 1
	expected := 1
	actual := alien.life

	assert.Equal(t, expected, actual)

	// Alien suddenly dies
	alien.Die()

	expected = 0
	actual = alien.life

	assert.Equal(t, expected, actual)
}
