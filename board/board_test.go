package board

import "testing"

func TestNew(t *testing.T) {
	const size = 9
	goban := make(GoBoard)
	goban.New(size)

	if goban.Size() != size {
		t.Errorf("Expected board size %d, but got %d", size, goban.Size())
	}

	for _, v := range goban {
		if v != Moku {
			t.Errorf("Expected all positions to be '%c', but found '%c'", Moku, v)
		}
	}
}
