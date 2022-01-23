package main

const (
	LIMIT_TOP_Y    = BORDER_SIZE*2 + BORDER_SPACING
	LIMIT_BOTTOM_Y = WINDOW_SIZE_Y - PADDLE_HEIGHT - LIMIT_TOP_Y
)

func yLimitOk(y float32) bool {
	return y <= LIMIT_BOTTOM_Y && y >= LIMIT_TOP_Y
}
