package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WINDOW_X = 800
	WINDOW_Y = 450
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
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()

		if rl.IsKeyDown(rl.KeyDown) {
			leftPaddle.move(MOVEMENT_SPEED)
		} else if rl.IsKeyDown(rl.KeyUp) {
			leftPaddle.move(-MOVEMENT_SPEED)
		}
	}

	rl.CloseWindow()
}
