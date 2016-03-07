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

var board Superboard

func NewGame() {
	board.Reset()

	curplayer = O
}

func Show() {
	board.Print()
	fmt.Println(curplayer, "to play")
}

// Play
// Plays a move at the given position
// Returns true if the game is over, false otherwise
func Play(pos position) bool {
	var player byte
	if curplayer == X {
		player = 'X'
	} else {
		player = 'O'
	}

	if board.Play(pos, player) {
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

// getPosFromInt
// Takes in an integer from 1-9
// Returns the equivalent position (think of a phone dial)
// Examples:
//	1 -> NW
//  6 -> E
func getPosFromInt(i int) position {
	var p position
	switch i {
	case 1:
		p = NW
	case 2:
		p = N
	case 3:
		p = NE
	case 4:
		p = W
	case 5:
		p = C
	case 6:
		p = E
	case 7:
		p = SW
	case 8:
		p = S
	case 9:
		p = SE
	}
	return p
}
