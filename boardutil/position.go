// Package boardutil contains utility function for working with a single Board
// position.go contains helper functions using the position enum (const in Golang)
package boardutil

type position int

// enum to store the 9 possible locations on the board
const (
	NW position = iota
	N
	NE
	W
	C
	E
	SW
	S
	SE
	NUM_SQUARES // = 9
)

// A map from strings to the corresponding position
var Position = map[string]position{
	"NW": NW,
	"N":  N,
	"NE": NE,
	"W":  W,
	"C":  C,
	"E":  E,
	"SW": SW,
	"S":  S,
	"SE": SE,
}
