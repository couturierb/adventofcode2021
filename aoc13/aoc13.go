package aoc13

import (
	"adventofcode2021/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const SIZE_X = 40
const SIZE_Y = 6

func Ex1() {
	input := utils.ReadFileSplitLineToArray("aoc13/aoc13.txt")
	paper := [SIZE_Y + 1][SIZE_X + 1]string{}

	for _, value := range input {
		x, _ := strconv.Atoi(strings.Split(value, ",")[0])
		y, _ := strconv.Atoi(strings.Split(value, ",")[1])
		x = searchPosition(x, SIZE_X)
		paper[y][x] = "#"
	}

	count := 0
	for _, line := range paper {
		for _, value := range line {
			if value == "#" {
				count++
			}
		}
	}
	fmt.Println(count)
}

func Ex2() {
	input := utils.ReadFileSplitLineToArray("aoc13/aoc13.txt")
	paper := [SIZE_Y + 1][SIZE_X + 1]string{}

	for i, line := range paper {
		for j, _ := range line {
			paper[i][j] = "."
		}
	}

	maxX := []int{655, 327, 163, 81, 40}
	maxY := []int{447, 223, 111, 55, 27, 13, 6}

	for _, value := range input {
		x, _ := strconv.Atoi(strings.Split(value, ",")[0])
		y, _ := strconv.Atoi(strings.Split(value, ",")[1])
		x = searchAllPosition(x, maxX)
		y = searchAllPosition(y, maxY)
		paper[y][x] = "#"
	}

	for _, value := range paper {
		fmt.Println(value)
	}
}

func searchPosition(value int, max int) int {
	if value < max {
		return value
	}

	return int(float64(max) - math.Abs(float64(max)-float64(value)))
}

// fold along x=655
// fold along y=447
// fold along x=327
// fold along y=223
// fold along x=163
// fold along y=111
// fold along x=81
// fold along y=55
// fold along x=40
// fold along y=27
// fold along y=13
// fold along y=6
func searchAllPosition(value int, maxX []int) int {
	for _, max := range maxX {
		if value > max {
			value = int(float64(max) - math.Abs(float64(max)-float64(value)))
		}
	}

	return value
}
