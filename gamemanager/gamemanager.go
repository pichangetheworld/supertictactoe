/*
 * The GameManager controls the flow of the game, such as:
 * - starting a new game,
 * - passing control to the active player, and
 * - determining when a player has won
 */
package gamemanager

import (
	"github.com/pichangetheworld/supertictactoe/boardutil"
)

var board boardutil.Board

func NewGame() {
	board.Reset()

	board.PrintBoard()

	board.Play("NW", 'O')
	board.PrintBoard()

	board.Play("E", 'X')
	board.PrintBoard()
	board.Play("E", 'O')
	board.PrintBoard()
}
