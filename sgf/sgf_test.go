package sgf

import (
	"goingo/board"
	"testing"
)

func TestParseSGF(t *testing.T) {
	sgfContent := "(;GM[1]FF[4]CA[UTF-8]AP[CGoban:3]ST[2]RU[Japanese]SZ[19]KM[7.50];B[pd];W[dd];B[pq])"
	goban, err := ParseSGF(sgfContent)
	if err != nil {
		t.Errorf("ParseSGF failed with error %v", err)
	}

	if goban.Size() != 19 {
		t.Errorf("Expected board size 19, but got %d", goban.Size())
	}

	expectedStones := map[board.Coord]rune{
		{X: 'P', Y: 4}:  board.Black,
		{X: 'D', Y: 4}:  board.White,
		{X: 'P', Y: 17}: board.Black,
	}

	for coord, expectedPlayer := range expectedStones {
		if goban[coord] != expectedPlayer {
			t.Errorf("Expected stone %c at %c%d, but got %c", expectedPlayer, coord.X, coord.Y, goban[coord])
		}
	}
}
