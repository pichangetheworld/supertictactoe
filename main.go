package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Main function
func main() {
	// Create a new game
	NewGame()

	// Show the board to the user
	Show()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your play (Ctrl+D to exit): ")
	for scanner.Scan() {
		// read one line from stdin
		line := scanner.Text()

		// parse arguments
		args := strings.Split(line, " ")

		// whether the arguments are valid so far
		valid := true

		// whether to play or not
		play := false

		if len(args) == 0 {
			valid = false
		}

		switch args[0] {
		case "help":
			// print help
			valid = false
		case "play":
			// play a move
			if len(args) < 3 {
				valid = false
			} else {
				play = true
			}
		case "show":
			Show()
		default:
			valid = false
		}

		// Have confirmed that there are at least 3 arguments
		if play {
			// convert the first arg (board #) to an integer, err if not int
			b, err := strconv.Atoi(args[1])
			// convert the second arg (position) to an integer, err if not int
			p, err2 := strconv.Atoi(args[2])
			if err == nil && err2 == nil && b > 0 && b <= 9 && p > 0 && p <= 9 {
				if Play(getPosFromInt(b), getPosFromInt(p)) {
					break
				}
			} else {
				valid = false
			}
		}

		if !valid {
			printPrompt()
		}

		fmt.Print("Enter your play (Ctrl+D to exit): ")
	}
	fmt.Println()
}

func printPrompt() {
	fmt.Println("Please enter a valid play")
	fmt.Println("Commands:")
	fmt.Println("\thelp - Show this help")
	fmt.Println("\tshow - Show the current board status")
	fmt.Println("\tplay x y - Play a move on board x at position y (1-9)")
}
