package rules

const (
	BOARD_WIDTH  = 100
	BOARD_HEIGHT = 100
)

type Board struct {
	Pieces [BOARD_WIDTH][BOARD_HEIGHT]bool
	next   *Board
}

func nNeighbours(board *Board, x, y int) int {
	lowerLimit := func(n, limit int) int {
		if n < limit {
			return limit
		}
		return n
	}

	upperLimit := func(n, limit int) int {
		if n > limit {
			return limit
		}
		return n
	}

	xLower := lowerLimit(x-1, 0)
	xUpper := upperLimit(x+1, BOARD_WIDTH-1)
	yLower := lowerLimit(y-1, 0)
	yUpper := upperLimit(y+1, BOARD_HEIGHT-1)

	n := 0
	for i := xLower; i <= xUpper; i++ {
		for j := yLower; j <= yUpper; j++ {
			if x == i && y == j {
				continue
			}

			if board.Pieces[i][j] {
				n++
			}
		}
	}

	return n
}

func NewBoard() *Board {
	var a, b Board
	a.next = &b
	b.next = &a
	return &a
}

func UpdateBoard(board *Board) *Board {
	next := board.next
	for x := 0; x < BOARD_WIDTH; x++ {
		for y := 0; y < BOARD_HEIGHT; y++ {
			n := nNeighbours(board, x, y)
			if board.Pieces[x][y] {
				// Under or over population
				if n < 2 || n > 3 {
					next.Pieces[x][y] = false
				} else {
					next.Pieces[x][y] = true
				}
			} else {
				// Reproduction
				if n == 3 {
					next.Pieces[x][y] = true
				} else {
					next.Pieces[x][y] = false
				}
			}
		}
	}

	return next
}
