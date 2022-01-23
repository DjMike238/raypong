package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	PADDLE_CENTER_Y = WINDOW_SIZE_Y/2 - PADDLE_HEIGHT/2

	PADDLE_WIDTH  = 8
	PADDLE_HEIGHT = 128

	MOVEMENT_SPEED = 4
)

type Paddle rl.Rectangle

var (
	leftPaddle = Paddle{
		X:      BORDER_SIZE*2 + BORDER_SPACING,
		Y:      PADDLE_CENTER_Y,
		Width:  PADDLE_WIDTH,
		Height: PADDLE_HEIGHT,
	}

	rightPaddle = Paddle{
		X:      WINDOW_SIZE_X - BORDER_SIZE*2 - BORDER_SPACING - PADDLE_WIDTH,
		Y:      PADDLE_CENTER_Y,
		Width:  PADDLE_WIDTH,
		Height: PADDLE_HEIGHT,
	}
)

func (p *Paddle) move(delta float32) {
	newY := p.Y + delta

	if yLimitOk(newY) {
		p.Y = newY
	}
}
