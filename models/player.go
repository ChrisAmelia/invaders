// Package models Player
package models

// Player represents the current player, the space vessel.
type Player struct {
	life      int
	x         int
	unicode   string
}

// NewPlayer returns an instance of a Player.
// Parameter life defines the Player's life.
// Parameter position defines the Player's abscissa.
// Parameter unicode defines the Player's string representation.
func NewPlayer(life int, position int, unicode string) Player {
	player := Player{life: life, x: position, unicode: unicode}

	return player
}

// DefaultPlayer returns an instance of a Player
// with default values set.
func DefaultPlayer() Player {
	player := NewPlayer(3, 0, "Λ")

	return player
}

// MoveRight moves the player one square to the right
// and returns its updated position.
func (player *Player) MoveRight() int {
	player.x += 1

	return player.x
}

// MoveLeft moves the player one square to the left
// and returns its updated position.
func (player *Player) MoveLeft() int {
	player.x -= 1

	return player.x
}

// GetHurt decreases the player's life by one.
func (player *Player) GetHurt() {
	player.life -= 1
}
