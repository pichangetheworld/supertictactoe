package main

import (
	"bufio"
	"fmt"
	"github.com/pichangetheworld/supertictactoe/gamemanager"
	"os"
)

func main() {
	gamemanager.NewGame()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text (Ctrl+D to exit): ")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Print("Enter text (Ctrl+D to exit): ")
	}
	fmt.Println()
}
