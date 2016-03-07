/*
 * The GameManager controls the flow of the game, such as:
 * - starting a new game,
 * - passing control to the active player, and
 * - determining when a player has won
 */
package main

import (
	"fmt"
)

// current player
const (
	O = iota
	X
)

var curplayer int

var board Board

func NewGame() {
	board.Reset()

	curplayer = O
}

func Show() {
	board.PrintBoard()
	fmt.Println(curplayer, "to play")
}

// Play
// Plays a move at the given position
// Returns true if the game is over, false otherwise
func Play(pos int) bool {
	var position string
	switch pos {
	case 1:
		position = "NW"
	case 2:
		position = "N"
	case 3:
		position = "NE"
	case 4:
		position = "W"
	case 5:
		position = "C"
	case 6:
		position = "E"
	case 7:
		position = "SW"
	case 8:
		position = "S"
	case 9:
		position = "SE"
	}
	var player byte
	if curplayer == X {
		player = 'X'
	} else {
		player = 'O'
	}

	if board.Play(position, player) {
		Show()
		if winner := board.Evaluate(); winner != ' ' {
			fmt.Println("Player", string(winner), "wins!")
			return true
		} else {
			curplayer = (curplayer + 1) % 2
		}
	}
	return false
}
