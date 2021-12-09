package aoc5

import (
	"adventofcode2021/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

}

func Ex1() {
	input := utils.ReadFileSplitLineToArray("aoc5/aoc5.txt")
	reg := regexp.MustCompile("[0-9]+,[0-9]+")
	results := [1000][1000]int{}

	for _, line := range input {
		values := reg.FindAllString(line, -1)
		firstValue := strings.Split(values[0], ",")
		secondValue := strings.Split(values[1], ",")
		x1, _ := strconv.Atoi(firstValue[0])
		y1, _ := strconv.Atoi(firstValue[1])
		x2, _ := strconv.Atoi(secondValue[0])
		y2, _ := strconv.Atoi(secondValue[1])

		// var minX int
		// var maxX int
		// if x2 > x1 {
		// 	minX = x1
		// 	maxX = x2
		// } else {
		// 	minX = x2
		// 	maxX = x1
		// }

		if x1 == x2 {
			if y2 > y1 {
				for y := y1; y <= y2; y++ {
					results[x1][y] += 1
				}
			} else {
				for y := y1; y >= y2; y-- {
					results[x1][y] += 1
				}
			}
		} else if y1 == y2 {
			if x2 > x1 {
				for x := x1; x <= x2; x++ {
					results[x][y1] += 1
				}
			} else {
				for x := x1; x >= x2; x-- {
					results[x][y1] += 1
				}
			}
		} else {
			if x2 > x1 && y2 > y1 {
				for x, y := x1, y1; x <= x2; x, y = x+1, y+1 {
					results[x][y] += 1
				}
			} else if x2 > x1 && y2 < y1 {
				for x, y := x1, y1; x <= x2; x, y = x+1, y-1 {
					results[x][y] += 1
				}
			} else if x2 < x1 && y2 > y1 {
				for x, y := x1, y1; x >= x2; x, y = x-1, y+1 {
					results[x][y] += 1
				}
			} else if x2 < x1 && y2 < y1 {
				for x, y := x1, y1; x >= x2; x, y = x-1, y-1 {
					results[x][y] += 1
				}
			}
		}
	}

	countTotal := 0
	for _, line := range results {
		for _, cellule := range line {
			if cellule >= 2 {
				countTotal += 1
			}
		}
	}

	fmt.Println("total : " + strconv.Itoa(countTotal))
	// for _, line := range results {
	// 	fmt.Println(line)
	// }

}
