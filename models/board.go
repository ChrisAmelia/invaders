package models

type Board struct {
	height, width int
	player Player
	aliens[] Alien
}

// NewBoard returns an instance of a board.
// The paramaeters width and height are speaking for themselves.
func NewBoard(width, height int) *Board {
	player := Player{
		life: 3,
		x: -1,
		unicode: "Λ",
	}

	board := new(Board)

	board.player = player
	board.width  = width
	board.height = height

	return board
}

// String returns a representation of the board.
// It is a map representing the game invaders
// and should be a square/rectangle map.
func (board Board) String() string {
	if board.width < 0 || board.height < 0 {
		panic("Dimensions cannot be negative.")
	}

	if board.width == 0 && board.height == 0 {
		return ""
	}

	// Draw the first line: +--------+
	mapString := "+"

	for j := 0 ; j < board.width ; j++ {
		mapString += "-"
	}

	mapString += "+\n"

	// Draw intermediate lines
	// Intermediate lines can either be empty: |        |
	// Or contain alien(s): |V   V  |
	for i := 0 ; i < board.height ; i++ {
		mapString += "|"

		for j := 0 ; j < board.width ; j++ {
			mapString += " "
		}

		mapString += "|\n"
	}

	// Draw the last line where the player is
	mapString += "*"

	for j := 0 ; j < board.width ; j++ {
		if j == board.player.x {
			mapString += board.player.unicode
		} else {
			mapString += " "
		}
	}

	mapString += "*"

	return mapString
}
