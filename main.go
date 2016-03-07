package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		// convert that line to an integer, err if not int
		if v, err := strconv.Atoi(line); err == nil {
			if v > 0 && v <= 9 {
				if Play(getPosFromInt(v)) {
					break
				}
			}
		} else {
			fmt.Print("Please enter a valid play")
		}
		fmt.Print("Enter your play (Ctrl+D to exit): ")
	}
	fmt.Println()
}
