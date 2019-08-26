package main

import "goingo/board"

const size = 9

var turn int
var goban map[coord]rune

func main() {
	goban = board.initBoard(size)
	board.printBoard(goban)
	/*moves = make(map[int]*move)
	turn++
	moves[turn].player = black
	moves[turn].lon = 'C'
	moves[turn].lat = 3*/
	//board.play(goban, moves[turn])
	//board.printBoard(goban)
}
