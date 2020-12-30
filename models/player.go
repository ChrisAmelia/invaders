// Package models Player
package models

// Player represents the current player, the space vessel.
type Player struct {
	life      uint
	x         int
	unicode   string
}

// NewPlayer returns an instance of a Player.
func NewPlayer(life uint, position int, unicode string) Player {
	player := Player{life: life, x: position, unicode: unicode}

	return player
}

// DefaultPlayer returns an instance of a Player
// with default values set.
func DefaultPlayer() Player {
	player := NewPlayer(3, 0, "Λ")

	return player
}
