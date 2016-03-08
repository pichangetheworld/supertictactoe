// Package boardutil contains utility function for working with a single Board
// position.go contains helper functions using the position enum (const in Golang)
package main

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
	NumSquares // = 9
)

// A list of winning combinations
// iterating through the list will determine whether teh board has been won
var Connections = [][]position{
	// horizontal
	{NW, N, NE},
	{W, C, E},
	{SW, S, SE},

	// vertical
	{NW, W, SW},
	{N, C, S},
	{NE, E, SE},

	// diagonal
	{NW, C, SE},
	{NE, C, SW},
}

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
