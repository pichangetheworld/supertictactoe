package main

import (
	"github.com/pichangetheworld/supertictactoe/boardutil"
)

func main() {
	board := boardutil.NewBoard()
	board.PrintBoard()

	board.Play("NW", 'O')
	board.PrintBoard()

	board.Play("E", 'X')
	board.PrintBoard()
	board.Play("E", 'O')
	board.PrintBoard()
}
