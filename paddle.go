package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	LIMIT_TOP_Y    = 20
	LIMIT_BOTTOM_Y = 366

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
		Y:      20,
		Width:  8,
		Height: 64,
	}

	rightPaddle = Paddle{
		X:      772,
		Y:      20,
		Width:  8,
		Height: 64,
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
