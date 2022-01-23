package main

import (
	"fmt"

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
)

var (
	playerPoints, cpuPoints int

	border = rl.Rectangle{
		X:      BORDER_SPACING,
		Y:      BORDER_SPACING,
		Width:  WINDOW_SIZE_X - BORDER_SPACING*2,
		Height: WINDOW_SIZE_Y - BORDER_SPACING*2,
	}

	lineStart = rl.Vector2{
		X: WINDOW_SIZE_X/2 - BORDER_SIZE/2 + BORDER_SPACING,
		Y: LIMIT_TOP_Y,
	}

	lineEnd = rl.Vector2{
		X: WINDOW_SIZE_X/2 - BORDER_SIZE/2 + BORDER_SPACING,
		Y: PADDLE_LIMIT_BOTTOM_Y + PADDLE_HEIGHT,
	}

	veryDarkGray = rl.Color{
		R: 20,
		G: 20,
		B: 20,
		A: 255,
	}
)

func main() {
	var gameStarted bool

	rl.InitWindow(WINDOW_SIZE_X, WINDOW_SIZE_Y, "RayPong by Dj_Mike238 - ESC to exit")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangleLinesEx(border, BORDER_SIZE, veryDarkGray)

		if gameStarted {
			movePaddle()

			rl.DrawRectangleRec(rl.Rectangle(leftPaddle), rl.LightGray)
			rl.DrawRectangleRec(rl.Rectangle(rightPaddle), rl.LightGray)
			rl.DrawLineEx(lineStart, lineEnd, BORDER_SIZE, veryDarkGray)
			rl.DrawRectangleRec(ball, rl.LightGray)

			drawScore()

		} else if rl.IsKeyPressed(rl.KeyEnter) {
			gameStarted = true

		} else {
			drawWelcomeText()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func drawScore() {
	player := fmt.Sprint(playerPoints)
	cpu := fmt.Sprint(cpuPoints)

	drawPoints(player, -SCORE_SPACING)
	drawPoints(cpu, SCORE_SPACING*2+BORDER_SIZE+1)
}

func drawPoints(score string, offset int32) {
	X := int32(lineStart.X) - rl.MeasureText(score, SCORE_FONT_SIZE) + offset
	rl.DrawText(score, X, LIMIT_TOP_Y, SCORE_FONT_SIZE, rl.LightGray)
}

func drawWelcomeText() {
	x := WINDOW_SIZE_X/2 - rl.MeasureText(WELCOME_TEXT, WELCOME_FONT_SIZE)/2
	rl.DrawText(WELCOME_TEXT, x, WELCOME_TEXT_Y, WELCOME_FONT_SIZE, rl.LightGray)
}

func movePaddle() {
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		leftPaddle.move(MOVEMENT_SPEED)
	} else if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		leftPaddle.move(-MOVEMENT_SPEED)
	}
}
