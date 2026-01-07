package sgf

import (
	"fmt"
	"goingo/board"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// SgfReader is a function to read a sgf file
func SgfReader(sgfFile string) {
	b, err := ioutil.ReadFile(sgfFile)
	if err != nil {
		log.Fatalf("failed to read sgf file: %v", err)
	}
	goban, err := ParseSGF(string(b))
	if err != nil {
		log.Fatalf("failed to parse sgf file: %v", err)
	}
	goban.Print()
}

// ParseSGF parses the content of an SGF file
func ParseSGF(sgfContent string) (board.GoBoard, error) {
	var goban board.GoBoard
	var size int

	// Extract board size
	re := regexp.MustCompile(`SZ\[(\d+)\]`)
	matches := re.FindStringSubmatch(sgfContent)
	if len(matches) > 1 {
		var err error
		size, err = strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse board size: %w", err)
		}
	} else {
		size = 19 // Default size
	}
	goban = make(board.GoBoard, size)
	goban.New(size)

	// Extract moves
	re = regexp.MustCompile(`;([BW])\[([a-z]{2})\]`)
	allMatches := re.FindAllStringSubmatch(sgfContent, -1)
	for _, move := range allMatches {
		var player rune
		if move[1] == "W" {
			player = board.White
		} else {
			player = board.Black
		}
		coord := board.Coord{
			X: rune(strings.ToUpper(string(move[2][0]))[0]),
			Y: int(move[2][1] - 'a' + 1),
		}
		goban.Play(coord, player)
	}
	return goban, nil
}
