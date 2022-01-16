package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WINDOW_X = 800
	WINDOW_Y = 450

	WELCOME_TEXT = " Welcome to RayPong!\nPress ENTER to begin."
	FONT_SIZE    = 20
)

func main() {
	rl.InitWindow(WINDOW_X, WINDOW_Y, "RayPong by Dj_Mike238")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangleLinesEx(frame, 8, rl.LightGray)
		rl.DrawRectangleRec(rl.Rectangle(leftPaddle), rl.LightGray)
		rl.DrawRectangleRec(rl.Rectangle(rightPaddle), rl.LightGray)
		rl.EndDrawing()

		movePaddle()
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
