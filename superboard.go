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
	{},
	{W, C, E},
	{},
	{SW, S, SE},
}

// Function to print out the full board in ASCII
func (sb *Superboard) Print() {
	for i := 0; i < numCols; i++ {
		if i != 0 {
			fmt.Println(boardSeparator)
		}

		board1, board2, board3 := sb.board[3*i], sb.board[3*i+1], sb.board[3*i+2]

		// iterate between each row of rows[] (see above)
		for j := 0; j < len(rows); j++ {
			var buffer bytes.Buffer

			// Print each individual board
			board1.PrintRow(j, &buffer)
			buffer.WriteString("‖")
			board2.PrintRow(j, &buffer)
			buffer.WriteString("‖")
			board3.PrintRow(j, &buffer)

			fmt.Println(buffer.String())
		}
	}
}

// Play
// Plays a move on board [b] by [player] at [position]
// Returns: whether that is a legal move or not
// e.g. sb.Play(E, N, "X")
//      sb.Play(N, N, "O")
//      sb.Play(N, N, "X") -> false
func (sb *Superboard) Play(b position, pos position, player byte) bool {
	return sb.board[b].Play(pos, player)
}

// Evaluate
// Determines if the game has been won
// Returns the character (e.g. 'O') if the player has won, or nil otherwise
func (sb *Superboard) Evaluate() byte {
	for _, conn := range Connections {
		//fmt.Println(conn)
		if p := sb.board[conn[0]].Evaluate(); p != ' ' &&
			sb.board[conn[1]].Evaluate() == p &&
			sb.board[conn[2]].Evaluate() == p {
			return p
		}
	}

	return ' '
}
