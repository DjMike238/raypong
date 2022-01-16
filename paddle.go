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
	frame = rl.Rectangle{
		X:      4,
		Y:      4,
		Width:  792,
		Height: 442,
	}

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

func limitsOk(y, delta float32) bool {
	newY := y + delta
	return newY <= LIMIT_BOTTOM_Y && newY >= LIMIT_TOP_Y
}

func (p *Paddle) move(delta float32) {
	if limitsOk(p.Y, delta) {
		p.Y += delta
	}
}
