package main

import (
	"bufio"
	"fmt"
	"github.com/pichangetheworld/supertictactoe/gamemanager"
	"os"
	"strconv"
)

func main() {
	gamemanager.NewGame()
	gamemanager.Show()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your play (Ctrl+D to exit): ")
	for scanner.Scan() {
		line := scanner.Text()
		if v, err := strconv.Atoi(line); err == nil {
			if v > 0 && v <= 9 {
				if gamemanager.Play(v) {
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
