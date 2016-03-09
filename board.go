// Package boardutil contains utility function for working with a single Board
package main

import (
	"bytes"
	"fmt"
)

// Function to hold the board state
// ' ' (space) -- indicates an empty square
// 'X' (capital x) -- indicates a square occupied by player1
// 'O' (capital o) -- indicates a square occupied by player2
type Board struct {
	// Each board is a 3x3 board
	state [int(NumSquares)]byte
}

// Function to clear the board to a new, empty 3x3 board
func (board *Board) Reset() {
	for i := 0; i < int(NumSquares); i++ {
		board.state[i] = byte(' ')
	}
}

// Incomplete board in ASCII
//    | X | O
// ---+---+---
//    |   |
// ---+---+---
//    |   | X

// Completed board in ASCII
//            ‖   | X | O ‖
//    /‾‾‾\   ‖---+---+---‖   \ V /
//   |  O  |  ‖   | X | X ‖   > X <
//    \___/   ‖---+---+---‖   / ^ \
//            ‖   |   | O ‖
// ===+===+===‖===+===+===‖===+===+===
//            ‖           ‖   | O |
//    \ V /   ‖   /‾‾‾\   ‖---+---+---
//    > X <   ‖  |  O  |  ‖   | O |
//    / ^ \   ‖   \___/   ‖---+---+---
//            ‖           ‖ O | X |
// ===+===+===‖===+===+===‖===+===+===
//            ‖   |   | O ‖
//    /‾‾‾\   ‖---+---+---‖   \ V /
//   |  O  |  ‖ O |   |   ‖   > X <
//    \___/   ‖---+---+---‖   / ^ \
//            ‖ X |   |   ‖

// Function to print out the board in ASCII
// PrintRow
// Prints out a row of a mini-board if it is incomplete,
// or a big symbol if the board has been won
// Parameters:
//    j      - integer from 1-5, indicates which row of the board to print
//    buffer - byteBuffer to write the string result into
func (board *Board) PrintRow(j int, buffer *bytes.Buffer) {

	switch board.Evaluate() {
	case ' ':
		if j%2 == 0 {
			// rows 0, 2, 4 are rows to print
			l, m, r := rows[j][0], rows[j][1], rows[j][2]
			buffer.WriteString(" ")
			buffer.WriteString(string(board.state[l]))
			buffer.WriteString(" | ")
			buffer.WriteString(string(board.state[m]))
			buffer.WriteString(" | ")
			buffer.WriteString(string(board.state[r]))
			buffer.WriteString(" ")
		} else {
			buffer.WriteString("---+---+---")
		}
	case 'X':
		buffer.WriteString(bigX[j])
	case 'O':
		buffer.WriteString(bigO[j])
	}
}

// Big X to draw over the board when X has won a mini-board
var bigX = []string{
	"           ",
	"   \\ V /   ",
	"   > X <   ",
	"   / ^ \\   ",
	"           ",
}

// Big O to draw over the board when O has won a mini-board
var bigO = []string{
	"           ",
	"   /‾‾‾\\   ",
	"  |  O  |  ",
	"   \\___/   ",
	"           ",
}

// Play
// Plays a move by [player] at [position]
// Returns: whether that is a legal move or not
// e.g. board.Play("N", "X")
//      board.Play("SW", "O")
//      board.Play("SW", "X") -> false

// point to *Board (not Board) so we actually change its underlying value
func (board *Board) Play(pos position, player byte) bool {
	if w := board.Evaluate(); w != ' ' {
		fmt.Println("Board is already won by ", w)
		return false
	}
	// check if the position is already taken
	if board.state[pos] != ' ' {
		fmt.Println("Position", pos, "was already taken")
		return false
	}
	board.state[pos] = player
	return true
}

// Evaluate
// Determines if this specific board has been won
// Returns the character (e.g. 'O') if the player has won, or nil otherwise
func (board *Board) Evaluate() byte {
	for _, conn := range Connections {
		//fmt.Println(conn)
		if p := board.state[conn[0]]; p != ' ' &&
			board.state[conn[1]] == p &&
			board.state[conn[2]] == p {
			return p
		}
	}

	return ' '
}
