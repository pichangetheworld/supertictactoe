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

// keep track of the last move so you can undo it,
// also so you can keep track of what is the next legal move
var lastBoard, lastMove position

func NewGame() {
	board.Reset()

	lastBoard = 0
	lastMove = 0
	curplayer = O
}

func Show() {
	board.Print()
	if lastMove != 0 {
		fmt.Println("Last move was play", int(lastBoard)+1, int(lastMove)+1)
	}
	fmt.Println(curplayer, "to play")
}

// Play
// Plays a move on board [b] at the given position [pos]
// Returns true if the game is over, false otherwise
func Play(b position, pos position) bool {
	var player byte
	if curplayer == X {
		player = 'X'
	} else {
		player = 'O'
	}

	if lastMove != 0 && b != lastMove {
		fmt.Println("Illegal move: You must play on board", int(lastMove)+1)
	} else if board.Play(b, pos, player) {
		Show()
		lastBoard, lastMove = b, pos
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
