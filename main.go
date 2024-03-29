package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_SIZE_X = 800
	WINDOW_SIZE_Y = 450

	WELCOME_TEXT      = " Welcome to RayPong!\nPress ENTER to begin."
	WELCOME_FONT_SIZE = 20
	WELCOME_TEXT_Y    = PADDLE_CENTER_Y + WELCOME_FONT_SIZE*2

	SCORE_FONT_SIZE = 50
	SCORE_SPACING   = 16

	BORDER_SIZE    = 8
	BORDER_SPACING = 4

	CENTER_LINE_X = WINDOW_SIZE_X/2 - BORDER_SIZE/2 + BORDER_SPACING
)

var (
	playerPoints int
	cpuPoints    int

	border = rl.Rectangle{
		X:      BORDER_SPACING,
		Y:      BORDER_SPACING,
		Width:  WINDOW_SIZE_X - BORDER_SPACING*2,
		Height: WINDOW_SIZE_Y - BORDER_SPACING*2,
	}

	lineStart = rl.Vector2{
		X: CENTER_LINE_X,
		Y: LIMIT_TOP_Y,
	}

	lineEnd = rl.Vector2{
		X: CENTER_LINE_X,
		Y: PADDLE_LIMIT_BOTTOM_Y + PADDLE_HEIGHT,
	}

	veryDarkGray = rl.Color{
		R: 20,
		G: 20,
		B: 20,
		A: 255,
	}

	leftPointRect = rl.Rectangle{
		X:      LIMIT_LEFT_X - 1,
		Y:      LIMIT_TOP_Y,
		Width:  PADDLE_WIDTH,
		Height: LIMIT_BOTTOM_Y - (BORDER_SIZE + BORDER_SPACING),
	}

	rightPointRect = rl.Rectangle{
		X:      LIMIT_RIGHT_X + 1,
		Y:      LIMIT_TOP_Y,
		Width:  PADDLE_WIDTH,
		Height: LIMIT_BOTTOM_Y - (BORDER_SIZE + BORDER_SPACING),
	}
)

func main() {
	var gameStarted bool

	rand.Seed(time.Now().UnixMilli())

	rl.InitWindow(WINDOW_SIZE_X, WINDOW_SIZE_Y, "RayPong by Dj_Mike238 - ESC to exit")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangleLinesEx(border, BORDER_SIZE, veryDarkGray)

		if gameStarted {
			rl.DrawRectangleRec(leftPointRect, rl.Black)
			rl.DrawRectangleRec(rightPointRect, rl.Black)

			moveLeftPaddle()
			moveRightPaddle()

			rl.DrawRectangleRec(rl.Rectangle(leftPaddle), rl.LightGray)
			rl.DrawRectangleRec(rl.Rectangle(rightPaddle), rl.LightGray)
			rl.DrawLineEx(lineStart, lineEnd, BORDER_SIZE, veryDarkGray)

			if !ball.moving {
				go ball.launch()
			} else {
				checkScore()
			}

			drawScore()

			rl.DrawRectangleRec(rl.Rectangle(ball.rect), ball.color)

		} else if rl.IsKeyPressed(rl.KeyEnter) {
			gameStarted = true

		} else {
			drawWelcomeText()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func checkScore() {
	if rl.CheckCollisionRecs(leftPointRect, ball.rect) {
		ball.stop <- true
		cpuPoints += 1

	} else if rl.CheckCollisionRecs(rightPointRect, ball.rect) {
		ball.stop <- true
		playerPoints += 1
	}
}

func drawScore() {
	player := fmt.Sprint(playerPoints)
	cpu := fmt.Sprint(cpuPoints)

	drawPoints(player, true)
	drawPoints(cpu, false)
}

func drawPoints(score string, isPlayer bool) {
	posX := int32(lineStart.X)

	if isPlayer {
		posX -= rl.MeasureText(score, SCORE_FONT_SIZE) + SCORE_SPACING
	} else {
		posX += SCORE_SPACING
	}

	rl.DrawText(score, posX, LIMIT_TOP_Y, SCORE_FONT_SIZE, veryDarkGray)
}

func drawWelcomeText() {
	X := WINDOW_SIZE_X/2 - rl.MeasureText(WELCOME_TEXT, WELCOME_FONT_SIZE)/2
	rl.DrawText(WELCOME_TEXT, X, WELCOME_TEXT_Y, WELCOME_FONT_SIZE, rl.LightGray)
}

func moveLeftPaddle() {
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		leftPaddle.move(MOVEMENT_SPEED)
	} else if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		leftPaddle.move(-MOVEMENT_SPEED)
	}
}

func moveRightPaddle() {
	if ball.moving && ball.rect.X >= CENTER_LINE_X {
		if ball.rect.Y > (rightPaddle.Y + PADDLE_HEIGHT) {
			rightPaddle.move(MOVEMENT_SPEED / 2)
		} else if ball.rect.Y < rightPaddle.Y {
			rightPaddle.move(-MOVEMENT_SPEED / 2)
		}
	}
}
