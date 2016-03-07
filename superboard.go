// superboard contains information about all 9 boards
package main

import (
	"bytes"
	"fmt"
)

type Superboard struct {
	board [int(NumSquares)]Board
}

func (sb *Superboard) Reset() {
	for i := 0; i < int(NumSquares); i++ {
		sb.board[i].Reset()
	}
}

// Number of rows and columns in a board
const numCols = 3

// Hardcoded horizontal ASCII separators
const lineSeparator = "---+---+---‖---+---+---‖---+---+---"
const boardSeparator = "===+===+===‖===+===+===‖===+===+==="

var rows = [][]position{
	{NW, N, NE},
	{W, C, E},
	{SW, S, SE},
}

// Function to print out the full board in ASCII
func (sb *Superboard) Print() {
	for i := 0; i < numCols; i++ {
		if i != 0 {
			fmt.Println(boardSeparator)
		}

		board1, board2, board3 := sb.board[i], sb.board[i+1], sb.board[i+2]

		for j := 0; j < len(rows); j++ {
			if j != 0 {
				fmt.Println(lineSeparator)
			}

			var buffer bytes.Buffer

			// left, middle, right
			l, m, r := rows[j][0], rows[j][1], rows[j][2]

			buffer.WriteString(" ")
			buffer.WriteString(string(board1.state[l]))
			buffer.WriteString(" | ")
			buffer.WriteString(string(board1.state[m]))
			buffer.WriteString(" | ")
			buffer.WriteString(string(board1.state[r]))
			buffer.WriteString(" ‖ ")

			buffer.WriteString(string(board2.state[l]))
			buffer.WriteString(" | ")
			buffer.WriteString(string(board2.state[m]))
			buffer.WriteString(" | ")
			buffer.WriteString(string(board2.state[r]))
			buffer.WriteString(" ‖ ")

			buffer.WriteString(string(board3.state[l]))
			buffer.WriteString(" | ")
			buffer.WriteString(string(board3.state[m]))
			buffer.WriteString(" | ")
			buffer.WriteString(string(board3.state[r]))

			fmt.Println(buffer.String())
		}
	}
	fmt.Println()
}

// Play
// Plays a move on board [b] by [player] at [position]
// Returns: whether that is a legal move or not
// e.g. sb.Play(E, N, "X")
//      sb.Play(N, N, "O")
//      sb.Play(N, N, "X") -> false
func (sb *Superboard) Play(pos position, player byte) bool {
	return true
}

// Evaluate
// Determines if the game has been won
// Returns the character (e.g. 'O') if the player has won, or nil otherwise
func (sb *Superboard) Evaluate() byte {
	return ' '
}
