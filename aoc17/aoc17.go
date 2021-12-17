package aoc17

import "fmt"

const MIN_X = 143
const MAX_X = 177
const MAX_Y = -71
const MIN_Y = -106

// Test
// const MIN_X = 20
// const MAX_X = 30
// const MAX_Y = -5
// const MIN_Y = -10

func Ex1() {
	test := 105
	fmt.Println(test * (test + 1) / 2)
}

func Ex2() {
	howMany := 0
	for x := 0; x <= MAX_X; x++ {
		for y := MIN_Y; y <= -MIN_Y; y++ {
			howMany += calcuHit(0, 0, x, y)
		}
	}
	fmt.Println(howMany)
}

func calcuHit(x_position int, y_position int, x_velocity int, y_velocity int) int {
	if y_position < MIN_Y {
		return 0
	}

	if x_position >= MIN_X && x_position <= MAX_X && y_position <= MAX_Y && y_position >= MIN_Y {
		return 1
	}

	new_x_position := x_position + x_velocity
	new_x_velocity := x_velocity - 1
	if new_x_velocity < 0 {
		new_x_velocity = 0
	}
	new_y_position := y_position + y_velocity
	new_y_velocity := y_velocity - 1
	return calcuHit(new_x_position, new_y_position, new_x_velocity, new_y_velocity)
}
