package main

import (
	"fmt"
	"gch/board"
)

func main() {
	board := board.NewGame()
	fmt.Printf("%v", board)
}
