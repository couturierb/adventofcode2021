package aoc15

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
)

func Ex1() {
	input := utils.ReadFileSplitLineToArray("aoc15/aoc15.txt")
	var grid [][]int

	for _, value := range input {
		grid = append(grid, utils.ConvertStringToIntArray(value))
	}

	total := 0
	parcours(grid, 0, 0, total)
	fmt.Println("total : " + strconv.Itoa(total))
}

func parcours(grid [][]int, x int, y int, total int) {
	scopedGrid := make([][]int, len(grid))
	copy(scopedGrid, grid)
	total += scopedGrid[y][x]
	scopedGrid[y][x] = 0

	fmt.Println()
	utils.PrintGrid(scopedGrid)

	// end
	if x == len(scopedGrid[0])-1 && y == len(scopedGrid)-1 {
		fmt.Println(total)
		return
	}

	// gauche
	if x > 0 && scopedGrid[y][x-1] != 0 {
		// totalLeft := make(chan int)
		parcours(scopedGrid, x-1, y, total)
		// total += <-totalLeft
	}

	// droite
	if x < len(scopedGrid[0])-1 && scopedGrid[y][x+1] != 0 {
		// totalRight := make(chan int)
		parcours(scopedGrid, x+1, y, total)
		// total += <-totalRight
	}

	// haut
	if y > 0 && scopedGrid[y-1][x] != 0 {
		// totalUp := make(chan int)
		parcours(scopedGrid, x, y-1, total)
		// total += <-totalUp
	}

	// bas
	if y < len(scopedGrid)-1 && scopedGrid[y+1][x] != 0 {
		// totalBottom := make(chan int)
		parcours(scopedGrid, x, y+1, total)
		// total += <-totalBottom
	}

	// channel <- total
}
