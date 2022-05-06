package main

import (
	"fmt"
	"gch/board"
)

func main() {
	board := board.NewGame()
	for {
		fmt.Printf("%v", board)
		fmt.Print("\nSelect Move: ")
		var move string
		fmt.Scan(&move)
		err := board.MovePieceFromString(move)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
