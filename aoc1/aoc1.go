package aoc1

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
)

func Ex1() {
	previous := 0
	total := 0
	numbers := utils.ReadFileSplitLineToArray("aoc1/aoc1.txt")

	for _, value := range numbers {
		currentMeasurment, _ := strconv.Atoi(value)
		if previous > 0 && currentMeasurment > previous {
			total += 1
		}
		previous = currentMeasurment
	}

	fmt.Println("total ex1 : ", total)
}

func Ex2() {
	total := 0
	numbers := utils.ReadFileSplitLineToArray("aoc1/aoc1.txt")

	for key, value := range numbers {
		currentMeasurment, _ := strconv.Atoi(value)
		if key < len(numbers)-3 {
			nextWindowMeasurment, _ := strconv.Atoi(numbers[key+3])
			if nextWindowMeasurment > currentMeasurment {
				total += 1
			}
		}
	}

	fmt.Println("total ex2 : ", total)
}
