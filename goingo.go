package main

import "goingo/board"

const size = 9

var turn int

func main() {
	goban := make(board.GoBoard, size)
	goban.New(size)
	goban.Print()
	/*moves = make(map[int]*move)
	turn++
	moves[turn].player = black
	moves[turn].lon = 'C'
	moves[turn].lat = 3*/
	//board.play(goban, moves[turn])
	//board.printBoard(goban)
}
