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

// Custom type for ball launch direction.
type Direction uint8

const (
	Random Direction = iota
	Left
	Right
)

type Ball struct {
	color     rl.Color
	moving    bool
	rect      rl.Rectangle
	direction Direction
	stop      chan bool
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
		deltaX = getDirection(b.direction)
		deltaY = getDirection(b.direction)
	)

	b.stop = make(chan bool, 1)

	for {
		select {
		case <-b.stop:
			close(b.stop)
			b.moving = false
			return

		default:
			x := b.rect.X + deltaX
			y := b.rect.Y + deltaY

			if !xLimitOk(x) || leftPaddle.collidesWith(*b) || rightPaddle.collidesWith(*b) {
				deltaX *= -1
			}

			b.rect.X += deltaX

			if !yLimitOk(y) {
				deltaY *= -1
			}

			b.rect.Y += deltaY

			time.Sleep(BALL_SPEED * time.Millisecond)
		}
	}
}

func getDirection(d Direction) float32 {
	switch d {
	case Random:
		d = Direction(rand.Intn(2) + 1)
	case Left:
		return -1
	case Right:
		return 1
	}

	return getDirection(d)
}
