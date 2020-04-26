package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	CANVAS_WIDTH = 1024
	CANVAS_HEIGHT = 1024

	BOARD_WIDTH = 100
	BOARD_HEIGHT = 100

	GRID_WIDTH = float64(CANVAS_WIDTH) / BOARD_WIDTH
	GRID_HEIGHT = float64(CANVAS_HEIGHT) / BOARD_HEIGHT
)

type Board [BOARD_WIDTH][BOARD_HEIGHT]bool

func getTitle(running bool) string {
	var status string
	if running {
		status = "Running"
	} else {
		status = "Paused"
	}

	return fmt.Sprintf("Game of Life: %s", status)
}

func initGrid() *imdraw.IMDraw {
	grid := imdraw.New(nil)
	grid.Color = colornames.Lightgrey

	for i := 1.0; i < BOARD_WIDTH; i++ {
		x := i * GRID_WIDTH
		grid.Push(pixel.V(x, 0), pixel.V(x, CANVAS_HEIGHT))
		grid.Line(1)
	}
	for i := 1.0; i < BOARD_HEIGHT; i++ {
		y := i * GRID_HEIGHT
		grid.Push(pixel.V(0, y), pixel.V(CANVAS_WIDTH, y))
		grid.Line(1)
	}

	return grid
}

func updatePieces(pieces *imdraw.IMDraw, board *Board) {
	pieces.Clear()
	pieces.Color = colornames.Yellow

	for x := 0; x < BOARD_WIDTH; x++ {
		for y := 0; y < BOARD_HEIGHT; y++ {
			if board[x][y] {
				xLower := float64(x)*GRID_WIDTH + 1
				xUpper := float64(x+1)*GRID_WIDTH - 1
				yLower := float64(y)*GRID_HEIGHT + 1
				yUpper := float64(y+1)*GRID_HEIGHT - 1

				pieces.Push(pixel.V(xLower, yLower),
					    pixel.V(xLower, yUpper),
					    pixel.V(xUpper, yUpper),
					    pixel.V(xUpper, yLower))
				pieces.Polygon(0)
			}
		}
	}
}

func getValueUnderMouse(win *pixelgl.Window, board *Board) *bool {
	pos := win.MousePosition()
	if pos.X < 0 || pos.X >= CANVAS_WIDTH ||
	   pos.Y < 0 || pos.Y >= CANVAS_HEIGHT {
		   return nil
        }


	x := int(pos.X / GRID_WIDTH)
	y := int(pos.Y / GRID_HEIGHT)

	return &board[x][y]
}

func run() {
	running := false

	cfg := pixelgl.WindowConfig {
		Title: getTitle(running),
		Bounds: pixel.R(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT),
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	grid := initGrid()

	var board Board

	pieces := imdraw.New(nil)
	updatePieces(pieces, &board)

	for !win.Closed() {
		if win.JustPressed(pixelgl.KeySpace) {
			running = !running
			win.SetTitle(getTitle(running))
		}
		if win.Pressed(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}
		if win.Pressed(pixelgl.MouseButtonLeft) {
			value := getValueUnderMouse(win, &board)

			if value != nil && !*value {
				*value = true
				updatePieces(pieces, &board)
			}
		}
		if win.Pressed(pixelgl.MouseButtonRight) {
			value := getValueUnderMouse(win, &board)

			if value != nil && *value {
				*value = false
				updatePieces(pieces, &board)
			}
		}

		win.Clear(colornames.Darkgrey)
		grid.Draw(win)
		pieces.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}