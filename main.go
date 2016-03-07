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

		if len(args) == 0 {
			valid = false
		}

		if valid {
			// convert that line to an integer, err if not int
			if v, err := strconv.Atoi(line); err == nil {
				if v > 0 && v <= 9 {
					if Play(getPosFromInt(v)) {
						break
					}
				} else {
					valid = false
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
}
