// Package alien Alien
package alien

import (
	"fmt"
)

// Alien represents the invader.
type Alien struct {
	life int
	x, y int
	unicode string
}

var (
	aliens []Alien
)

// ResetStateOfAliens sets the aliens to nil.
// For testing purpose.
func ResetStateOfAliens() {
	aliens = nil
}

func GetAliens() *[]Alien {
	return &aliens
}

// NewAlien returns an instance of Alien.
// Parameter life defines the life of the Alien.
// Parameter x defines the Alien's abscissa.
// Parameter y defines the Alien's ordinate.
// Parameter unicode defines the Alien's string representation.
func NewAlien(life, x, y int, unicode string) Alien {
	alien := Alien{
		life: life,
		x: x,
		y: y, 
		unicode: unicode,
	}

	return alien
}

// GetAlienAt returns the alien located at (x,y).
// If no alien is located at (x,y) then return a default alien and a non-nil error.
// Parameter x defines the Alien's abscissa.
// Parameter y defines the Alien's ordinate.
func GetAlienAt(x, y int) (Alien, error) {
	for _, alien := range aliens {
		if alien.x == x && alien.y == y {
			return alien, nil
		}
	}

	alien := Alien{life: -1, unicode: "ø"}

	return alien, fmt.Errorf("no alien found at (%d,%d)", x, y)
}

// DefaultAlien returns an instance of Alien
// with default values set: 1 life and a string representation.
// If an alien is already present at given coordinates then
// this alien is returned.
// Paramater x defines the Alien's abscissa.
// Parameter y defines the Alien's ordinate.
func DefaultAlien(x, y int) Alien {
	for _, alien := range aliens {
		if alien.x == x && alien.y == y {
			return alien
		}
	}

	alien := NewAlien(1, x, y, "V")

	aliens = append(aliens, alien)

	return alien
}

// Life returns the alien's life.
func (alien *Alien) Life() int {
	return alien.life
}

// SetLife sets the alien's life with given value.
func (alien *Alien) SetLife(life int) {
	alien.life = life
}

// X returns the alien's abscissa.
func (alien *Alien) X() int {
	return alien.x
}

// SetX sets the alien's abscissa life with given value.
func (alien *Alien) SetX(x int)  {
	alien.x = x
}

// Y returns the alien's ordinate.
func (alien *Alien) Y() int {
	return alien.y
}

// SetY sets the alien's ordinate with given value.
func (alien *Alien) SetY(y int) {
	alien.y = y
}

// Unicode returns the alien's unicode.
func (alien *Alien) Unicode() string {
	return alien.unicode
}

// SetUnicode sets the alien's life with given value.
func (alien *Alien) SetUnicode(unicode string) {
	alien.unicode = unicode
}

// MoveDown moves the alien one square down
// and returns its updated position.
func (alien *Alien) MoveDown() int {
	alien.y -= 1

	return alien.y
}

// GetHit decreases the alien's life by one.
func (alien *Alien) GetHit() {
	alien.life -= 1
}

// IsDead returns true if alien's life has reached zero,
// else returns false.
func (alien *Alien) IsDead() bool {
	return alien.life <= 0
}

// ChangeIconToEmpty sets the alien's unicode to an empty one if the alien is dead.
func (alien *Alien) ChangeIconToEmpty() {
	if (alien.IsDead()) {
		alien.unicode = ""
	}
}

// Die sets the alien's life to zero.
func (alien *Alien) Die() {
	alien.life = 0
}
