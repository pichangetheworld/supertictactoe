// Package boardutil contains utility function for working with a single Board
package boardutil

import (
	"fmt"
)

// Function to hold the board state
// ' ' (space) -- indicates an empty square
// 'X' (capital x) -- indicates a square occupied by player1
// 'O' (capital o) -- indicates a square occupied by player2
type Board struct {
	// Each board is a 3x3 board
	state [int(NUM_SQUARES)]byte
}

// Function to clear the board to a new, empty 3x3 board
func (board *Board) Reset() {
	for i := 0; i < int(NUM_SQUARES); i++ {
		board.state[i] = byte(' ')
	}
}

// Function to print out the board in ASCII
// Use a method pointer (*Board) instead of (Board) so we can just pass the pointer around
// This is faster than creating a copy every time
func (board *Board) PrintBoard() {
	fmt.Printf(" %s | %s | %s \n", string(board.state[NW]), string(board.state[N]), string(board.state[NE]))
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s \n", string(board.state[W]), string(board.state[C]), string(board.state[E]))
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s \n", string(board.state[SW]), string(board.state[S]), string(board.state[SE]))
	fmt.Println()
}

// Plays a move by [player] at [position]
// Returns: whether that is a legal move or not
// e.g. board.Play("N", "X")
//      board.Play("SW", "O")
//      board.Play("SW", "X") -> false

// point to *Board (not Board) so we actually change its underlying value
func (board *Board) Play(pos string, player byte) bool {
	// TODO: check if the position is already taken
	if board.state[Position[pos]] != ' ' {
		fmt.Println("Position", pos, "was already taken")
		return false
	}
	board.state[Position[pos]] = player
	return true
}
