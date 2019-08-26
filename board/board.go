package board

import "fmt"

const moku = '+'
const white = 'w'
const black = 'b'
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const maxSize = 19

//coord is a structure for coordinates on the go board.
type coord struct {
	col rune //column letter
	lin int  //line number
}

//move is a structure to store moves data : coordinates and player.
type move struct {
	player rune //player color
	loc    coord
}

//initBoard initializes an empty board of size "size".
func initBoard(size int) map[coord]rune {
	//func initBoard(size int) [][]rune {
	if size > maxSize {
		fmt.Println("Size > 19. The board will default to 19*19.")
		size = maxSize
	}
	/*goban := make([][]rune, size)
	for i := 0; i < len(goban); i++ {
		for j := 0; j < len(goban); j++ {
			goban[i][j] = moku
		}
	}*/
	goban := make(map[coord]rune)
	for c, le := range alphabet {
		if c == size {
			break
		}
		for li := 1; li < size+1; li++ {
			goban[coord{le, li}] = moku
		}
	}
	return goban
}

//printBoard prints the board on stdout.
func printBoard(goban map[coord]rune) {
	//func printBoard(goban [][]rune) {
	/*fmt.Println("   A  B  C  D  E  F  G  H  I")
	for i := 0; i < len(goban); i++ {
		fmt.Print(i+1, " ")
		for j := 0; j < len(goban); j++ {
			fmt.Printf(" %c ", goban[i][j])
		}
		fmt.Print("\n")
	}*/
	fmt.Print("   ")
	for li := 0; li < len(goban); li++ {
		if li > 0 {
			fmt.Print(li, " ")
		}
		for c, le := range alphabet {
			if li == 0 {
				fmt.Print(le, " ")
			}
			if c == len(goban) {
				break
			}
			fmt.Printf(" %c ", goban[coord{le, li}])
		}
	}
}

//play registers a move on the goban
/*func play(goban [][]rune, m move) [][]rune {
	goban[m.loc.lat][m.loc.lon] = m.player
	return goban
}*/
