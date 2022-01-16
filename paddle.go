package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	LIMIT_TOP_Y    = FRAME_BORDER_SIZE*2 + FRAME_SPACING
	LIMIT_BOTTOM_Y = WINDOW_SIZE_Y - PADDLE_HEIGHT - LIMIT_TOP_Y

	CENTER_Y = WINDOW_SIZE_Y/2 - PADDLE_HEIGHT/2

	PADDLE_WIDTH  = 8
	PADDLE_HEIGHT = 64

	MOVEMENT_SPEED = 4
)

type Paddle rl.Rectangle

var (
	leftPaddle = Paddle{
		X:      FRAME_BORDER_SIZE*2 + FRAME_SPACING,
		Y:      CENTER_Y,
		Width:  PADDLE_WIDTH,
		Height: PADDLE_HEIGHT,
	}

	rightPaddle = Paddle{
		X:      WINDOW_SIZE_X - FRAME_BORDER_SIZE*2 - FRAME_SPACING - PADDLE_WIDTH,
		Y:      CENTER_Y,
		Width:  PADDLE_WIDTH,
		Height: PADDLE_HEIGHT,
	}
)

func limitsOk(y float32) bool {
	return y <= LIMIT_BOTTOM_Y && y >= LIMIT_TOP_Y
}

func (p *Paddle) move(delta float32) {
	newY := p.Y + delta

	if limitsOk(newY) {
		p.Y = newY
	}
}
