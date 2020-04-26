package rules

import (
	"math/rand"
)

const (
	BOARD_WIDTH  = 100
	BOARD_HEIGHT = 100
)

type Board [BOARD_WIDTH][BOARD_HEIGHT]bool

func UpdateBoard(board *Board) {
	xRand := rand.Intn(BOARD_WIDTH)
	yRand := rand.Intn(BOARD_HEIGHT)

	board[xRand][yRand] = !board[xRand][yRand]
}
