package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WINDOW_SIZE_X = 800
	WINDOW_SIZE_Y = 450

	WELCOME_TEXT = " Welcome to RayPong!\nPress ENTER to begin."
	FONT_SIZE    = 20
)

var (
	gameStarted bool
)

func main() {
	rl.InitWindow(WINDOW_SIZE_X, WINDOW_SIZE_Y, "RayPong by Dj_Mike238")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangleLinesEx(frame, 8, rl.LightGray)
		rl.DrawRectangleRec(rl.Rectangle(leftPaddle), rl.LightGray)
		rl.DrawRectangleRec(rl.Rectangle(rightPaddle), rl.LightGray)
		rl.EndDrawing()

		if gameStarted {
			movePaddle()
			rl.DrawRectangleRec(rl.Rectangle(ball), rl.LightGray)
		} else if rl.IsKeyPressed(rl.KeyEnter) {
			gameStarted = true
		} else {
			drawText()
		}
	}

	rl.CloseWindow()
}

func movePaddle() {
	if rl.IsKeyDown(rl.KeyDown) {
		leftPaddle.move(MOVEMENT_SPEED)
	} else if rl.IsKeyDown(rl.KeyUp) {
		leftPaddle.move(-MOVEMENT_SPEED)
	}
}

func drawText() {
	x := rl.MeasureText(WELCOME_TEXT, FONT_SIZE)
	rl.DrawText(WELCOME_TEXT, x, 200, FONT_SIZE, rl.LightGray)
}
