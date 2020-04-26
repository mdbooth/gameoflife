package rules

import (
	"math/rand"
)

const (
	BOARD_WIDTH  = 100
	BOARD_HEIGHT = 100
)

type Board struct {
	Pieces [BOARD_WIDTH][BOARD_HEIGHT]bool
	next   *Board
}

func NewBoard() *Board {
	var a, b Board
	a.next = &b
	b.next = &a
	return &a
}

func UpdateBoard(board *Board) *Board {
	board.next.Pieces = board.Pieces

	xRand := rand.Intn(BOARD_WIDTH)
	yRand := rand.Intn(BOARD_HEIGHT)

	board.next.Pieces[xRand][yRand] = true
	return board.next
}
