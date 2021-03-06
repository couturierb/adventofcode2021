package aoc9

import (
	"adventofcode2021/utils"
	"fmt"
)

func Ex1() {
	input := utils.ReadFileSplitLineToArray("aoc9/aoc9.txt")

	var grid [][]int
	for _, value := range input {
		grid = append(grid, utils.ConvertStringToIntArray(value))
	}

	var low_points []int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if isLowPoint(grid, x, y) {
				low_points = append(low_points, grid[x][y])
			}
		}
	}

	var total int
	for _, value := range low_points {
		total += value + 1
	}
	fmt.Println(total)
}

func isLowPoint(grid [][]int, x int, y int) bool {
	current_value := grid[x][y]

	if x > 0 && grid[x-1][y] <= current_value {
		return false
	}

	if x < len(grid)-1 && grid[x+1][y] <= current_value {
		return false
	}

	if y > 0 && grid[x][y-1] <= current_value {
		return false
	}

	if y < len(grid[0])-1 && grid[x][y+1] <= current_value {
		return false
	}

	return true
}

func Ex2() {
	input := utils.ReadFileSplitLineToArray("aoc9/aoc9.txt")

	var grid [][]int
	for _, value := range input {
		grid = append(grid, utils.ConvertStringToIntArray(value))
	}

	var basins []int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if isLowPoint(grid, x, y) {
				basinFounded := countBasinSize(&grid, x, y)
				basins = append(basins, basinFounded)
			}
		}
	}

	basins, value1 := popGreatest(basins)
	basins, value2 := popGreatest(basins)
	basins, value3 := popGreatest(basins)
	fmt.Println(value1 * value2 * value3)
}

func countBasinSize(grid *[][]int, x int, y int) int {
	current_value := (*grid)[x][y]
	(*grid)[x][y] = 9
	totalBasin := 1

	if current_value == 8 {
		return totalBasin
	}

	// haut
	if x > 0 && (*grid)[x-1][y] == current_value+1 {
		totalBasin += countBasinSize(grid, x-1, y)
	}

	// bas
	if x < len(*grid)-1 && (*grid)[x+1][y] == current_value+1 {
		totalBasin += countBasinSize(grid, x+1, y)
	}

	// gauche
	if y > 0 && (*grid)[x][y-1] == current_value+1 {
		totalBasin += countBasinSize(grid, x, y-1)
	}

	// droite
	if y < len((*grid)[0])-1 && (*grid)[x][y+1] == current_value+1 {
		totalBasin += countBasinSize(grid, x, y+1)
	}

	return totalBasin
}

func popGreatest(basins []int) ([]int, int) {
	var greatest int
	var greatestIndex int
	for index, value := range basins {
		if value > greatest {
			greatest = value
			greatestIndex = index
		}
	}

	return append(basins[:greatestIndex], basins[greatestIndex+1:]...), greatest
}
