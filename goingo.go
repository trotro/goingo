package main

import "goingo/board"

const size = 9

// Pos is used to represent a new move
type Pos struct {
	Player rune //player color
	Loc    board.Coord
}

var turn int

func main() {
	goban := make(board.GoBoard, size)
	goban.New(size)
	goban.Print()
	//moves := make(map[int]*Move)
	moves := make(map[int]*Pos)
	turn++
	curMove := new(Pos)
	curMove.Player = board.Black
	curMove.Loc = board.Coord{X: 'C', Y: 3}
	moves[turn] = curMove
	goban.Play(moves[turn].Loc, curMove.Player)
	goban.Print()
}
