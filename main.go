package main

import (
	"fmt"
)

func main() {
	var board [3][3]string
	printBoard(board)
}

func printBoard(board [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(" | ")
			fmt.Print(board[i][j])
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("---------")
		}
	}
}
