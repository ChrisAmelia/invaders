package models

// Alien represents the invader.
type Alien struct {
	life int
	x, y int
	unicode string
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

// DefaultAlien returns an instance of Alien
// with default values set: 1 life and a string representation.
// Paramater x defines the Alien's abscissa.
// Parameter y defines the Alien's ordinate.
func DefaultAlien(x, y int) Alien {
	alien := NewAlien(1, x, y, "V")

	return alien
}

// MoveDown moves the alien one square down
// and returns its updated position.
func (alien *Alien) MoveDown() int {
	alien.y -= 1

	return alien.y
}
