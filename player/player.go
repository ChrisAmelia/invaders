// Package player Player
package player

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

// Life returns the player's life.
func (player *Player) Life() int {
	return player.life
}

// SetLife sets the player's life with given value.
func (player *Player) SetLife(life int) {
	player.life = life
}


// X returns the player's abscissa.
func (player *Player) X() int {
	return player.x
}

// SetX sets the player's abscissa with given value.
func (player *Player) SetX(x int) {
	player.x = x
}

// Unicode returns the player's unicode.
func (player *Player) Unicode() string {
	return player.unicode
}

// SetUnicode sets the player's unicode with given value.
func (player *Player) SetUnicode(unicode string)  {
	player.unicode = unicode
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

// GetHit decreases the player's life by one.
func (player *Player) GetHit() {
	player.life -= 1
}

// IsDead returns true if player's life has reached zero,
// else returns false.
func (player *Player) IsDead() bool {
	return player.life <= 0
}

// Die sets the player's life to zero.
func (player *Player) Die() {
	player.life = 0
}

// ChangeIconToSkull sets the player's unicode to a skull one if player is dead.
// See unicode 2620 (☠).
func (player *Player) ChangeIconToSkull() {
	if (player.IsDead()) {
		player.unicode = "☠"
	}
}
