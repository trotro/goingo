package main

import "fmt"

const size = 9
const moku = '+'
const white = 'w'
const black = 'b'

var board [size][size]rune

func init() {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			board[i][j] = moku
		}
	}
}

func printBoard() {
	fmt.Println("   A  B  C  D  E  F  G  H  I")
	for i := 0; i < len(board); i++ {
		fmt.Print(i+1, " ")
		for j := 0; j < len(board); j++ {
			fmt.Printf(" %c ", board[i][j])
		}
		fmt.Print("\n")
	}
}

func main() {
	printBoard()
}
