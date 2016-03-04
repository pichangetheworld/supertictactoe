// Package boardutil contains utility function for working with a single Board
package boardutil

import (
	"fmt"
)

type Board struct {
	// Each board is a 3x3 board
	NW byte
	N  byte
	NE byte
	W  byte
	C  byte
	E  byte
	SW byte
	S  byte
	SE byte
}

func NewBoard() Board {
	var board Board
	board.NW = byte(' ')
	board.N = byte(' ')
	board.NE = byte(' ')
	board.W = byte(' ')
	board.C = byte(' ')
	board.E = byte(' ')
	board.SW = byte(' ')
	board.S = byte(' ')
	board.SE = byte(' ')
	return board
}

func (board Board) PrintBoard() {
	fmt.Printf(" %s | %s | %s \n", string(board.NW), string(board.N), string(board.NE))
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s \n", string(board.W), string(board.C), string(board.E))
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s \n", string(board.SW), string(board.S), string(board.SE))
	fmt.Println()
}

// Plays a move by [player] at [position]
// Returns: whether that is a legal move or not
// e.g. board.Play("N", "X")
//      board.Play("SW", "O")
//      board.Play("SW", "X") -> false

// point to *Board (not Board) so we actually change its underlying value
func (board *Board) Play(position string, player byte) bool {
	// TODO: make position into an enum
	switch position {
	case "NW":
		// TODO: check if the position is already taken
		board.NW = player
	case "N":
		board.N = player
	case "E":
		board.E = player
	}
	return true
}
