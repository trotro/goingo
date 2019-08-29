package board

import (
	"fmt"
	"math"
)

const (
	//Moku is the rune to represent an empty moku on the go board
	Moku = '+'
	//Black is the rune to represent black stones
	Black = 'b'
	//White is the rune to represent white stones
	White = 'w'
	//unexported constants
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	maxSize  = 19
)

//Coord is a structure for coordinates on the go board.
type Coord struct {
	Col rune //column letter
	Lin int  //line number
}

//Move is a structure to store moves data : coordinates and player.
type Move struct {
	Player rune //player color
	Loc    Coord
}

//GoBoard is a map to represent the go board.
type GoBoard map[Coord]rune

//New initializes an empty board of size "size".
func (goban GoBoard) New(size int) {
	if size > maxSize {
		fmt.Println("Size > 19. The board will default to 19*19.")
		size = maxSize
	}
	for c, le := range alphabet {
		if c >= size {
			break
		}
		for li := 1; li < size+1; li++ {
			goban[Coord{le, li}] = Moku
		}
	}
}

//Size compute the size of the go board like humans do.
func (goban GoBoard) Size() int {
	return int(math.Sqrt(float64(len(goban))))
}

//Play registers a move on the goban
func (goban GoBoard) Play(m Move) {
	goban[Coord{m.Loc.Col, m.Loc.Lin}] = m.Player
}

//Print prints the board on stdout.
func (goban GoBoard) Print() {
	size := goban.Size()
	//fmt.Printf("Goban size : %v \n", size)
	fmt.Print("   ")
	for li := 0; li < size+1; li++ {
		if li > 0 {
			fmt.Print(li, " ")
		}
		for c, le := range alphabet {
			if li == 0 && c < size {
				fmt.Printf("%c", le)
			}
			if c >= size {
				break
			}
			fmt.Printf(" %c  ", goban[Coord{le, li}])
		}
		fmt.Println()
	}
}
