package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WINDOW_SIZE_X = 800
	WINDOW_SIZE_Y = 450

	WELCOME_TEXT = " Welcome to RayPong!\nPress ENTER to begin."
	FONT_SIZE    = 20

	FRAME_BORDER_SIZE = 8
	FRAME_SPACING     = 4
)

var (
	frame = rl.Rectangle{
		X:      FRAME_SPACING,
		Y:      FRAME_SPACING,
		Width:  WINDOW_SIZE_X - FRAME_SPACING*2,
		Height: WINDOW_SIZE_Y - FRAME_SPACING*2,
	}
)

func main() {
	var gameStarted bool

	rl.InitWindow(WINDOW_SIZE_X, WINDOW_SIZE_Y, "RayPong by Dj_Mike238 - ESC to exit")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangleLinesEx(frame, FRAME_BORDER_SIZE, rl.LightGray)

		if gameStarted {
			movePaddle()

			rl.DrawRectangleRec(rl.Rectangle(leftPaddle), rl.LightGray)
			rl.DrawRectangleRec(rl.Rectangle(rightPaddle), rl.LightGray)
			rl.DrawRectangleRec(ball, rl.LightGray)

		} else if rl.IsKeyPressed(rl.KeyEnter) {
			gameStarted = true

		} else {
			drawText()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func drawText() {
	x := WINDOW_SIZE_X/2 - rl.MeasureText(WELCOME_TEXT, FONT_SIZE)/2
	rl.DrawText(WELCOME_TEXT, x, 200, FONT_SIZE, rl.LightGray)
}

func movePaddle() {
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		leftPaddle.move(MOVEMENT_SPEED)
	} else if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		leftPaddle.move(-MOVEMENT_SPEED)
	}
}
