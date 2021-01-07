package models

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
