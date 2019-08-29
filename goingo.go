package main

import "goingo/board"

const size = 9

var turn int

func main() {
	goban := make(board.GoBoard, size)
	goban.New(size)
	goban.Print()
	moves := make(map[int]*board.Move)
	turn++
	curMove := new(board.Move)
	curMove.Player = board.Black
	curMove.Loc = board.Coord{Col: 'C', Lin: 3}
	moves[turn] = curMove
	goban.Play(*moves[turn])
	goban.Print()
}
