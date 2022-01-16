package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	ball = rl.Rectangle{
		X:      WINDOW_SIZE_X/2 - PADDLE_WIDTH/2,
		Y:      WINDOW_SIZE_Y/2 - PADDLE_WIDTH/2,
		Width:  PADDLE_WIDTH,
		Height: PADDLE_WIDTH,
	}
)
