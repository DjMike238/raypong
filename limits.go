package main

const (
	LIMIT_LEFT_X  = BORDER_SIZE*2 + BORDER_SPACING
	LIMIT_RIGHT_X = WINDOW_SIZE_X - LIMIT_LEFT_X - BORDER_SIZE

	LIMIT_TOP_Y    = BORDER_SIZE*2 + BORDER_SPACING
	LIMIT_BOTTOM_Y = WINDOW_SIZE_Y - LIMIT_TOP_Y - BORDER_SIZE

	PADDLE_LIMIT_TOP_Y    = BORDER_SIZE*2 + BORDER_SPACING
	PADDLE_LIMIT_BOTTOM_Y = WINDOW_SIZE_Y - PADDLE_HEIGHT - PADDLE_LIMIT_TOP_Y
)

func xLimitOk(x float32) bool {
	return x <= LIMIT_RIGHT_X && x >= LIMIT_LEFT_X
}

func yLimitOk(y float32) bool {
	return y <= LIMIT_BOTTOM_Y && y >= LIMIT_TOP_Y
}

func paddleLimitOk(y float32) bool {
	return y <= PADDLE_LIMIT_BOTTOM_Y && y >= PADDLE_LIMIT_TOP_Y
}
