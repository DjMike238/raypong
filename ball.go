package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// The lower it is, the faster the ball will move.
	BALL_SPEED = 8
)

type Ball struct {
	rect   rl.Rectangle
	moving bool
}

var (
	ball = Ball{
		rect: rl.Rectangle{
			Width:  PADDLE_WIDTH,
			Height: PADDLE_WIDTH,
		},
	}
)

func (b *Ball) launch() {
	pos := rand.Intn(LIMIT_BOTTOM_Y-LIMIT_TOP_Y) + LIMIT_TOP_Y
	b.rect.X = WINDOW_SIZE_X/2 - PADDLE_WIDTH/2
	b.rect.Y = float32(pos)
	b.moving = true
}

func (b *Ball) move() {
	var (
		deltaX = float32(1)
		deltaY = float32(1)

		ticker = time.NewTicker(BALL_SPEED * time.Millisecond)
	)

	for range ticker.C {
		x := b.rect.X + deltaX
		y := b.rect.Y + deltaY

		if !xLimitOk(x) {
			deltaX *= -1
		}

		b.rect.X += deltaX

		if !yLimitOk(y) {
			deltaY *= -1
		}

		b.rect.Y += deltaY

		if !b.moving {
			ticker.Stop()
		}
	}

	b.moving = false
}
