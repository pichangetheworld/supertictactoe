// superboard contains information about all 9 boards
package main

import (
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

// Number of columns in a board
const numCols = 3

// Function to print out the full board in ASCII
func (sb *Superboard) Print() {
	for i := 0; i < numCols; i++ {
		board1 := sb.board[i]
		//board1, board2, board3 := sb.board[i], sb.board[i+1], sb.board[i+2]
		fmt.Printf(" %s | %s | %s ‖\n", string(board1.state[NW]), string(board1.state[N]), string(board1.state[NE]))
		fmt.Println("---+---+---‖")
		fmt.Printf(" %s | %s | %s ‖\n", string(board1.state[W]), string(board1.state[C]), string(board1.state[E]))
		fmt.Println("---+---+---‖")
		fmt.Printf(" %s | %s | %s ‖\n", string(board1.state[SW]), string(board1.state[S]), string(board1.state[SE]))
		fmt.Println("===+===+===‖")
	}
	fmt.Println()
}

func (sb *Superboard) Play(pos string, player byte) bool {
	return true
}

func (sb *Superboard) Evaluate() byte {
	return ' '
}
