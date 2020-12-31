package models

// Alien represents the invader.
type Alien struct {
	life int
	x, y int
	unicode string
}

// NewAlien returns an instance of Alien.
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
// with default values set.
func DefaultAlien() Alien {
	alien := NewAlien(1, 0, 0, "V")

	return alien
}

// MoveDown moves the alien one square down
// and returns its updated position.
func (alien *Alien) MoveDown() int {
	alien.y -= 1

	return alien.y
}
