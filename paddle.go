package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	LIMIT_TOP_Y    = 20
	LIMIT_BOTTOM_Y = 366

	CENTER_Y = 192

	PADDLE_WIDTH  = 8
	PADDLE_HEIGHT = 64

	MOVEMENT_SPEED = 4
)

type Paddle rl.Rectangle

var (
	leftPaddle = Paddle{
		X:      20,
		Y:      CENTER_Y,
		Width:  PADDLE_WIDTH,
		Height: PADDLE_HEIGHT,
	}

	rightPaddle = Paddle{
		X:      772,
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
