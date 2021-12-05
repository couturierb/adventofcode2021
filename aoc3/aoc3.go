package aoc3

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
)

var LINE_SIZE int = 12

func Ex1() {
	input := utils.ReadFileSplitToArray("aoc3/aoc3.txt")

	countArray := make([]int, LINE_SIZE)
	for _, inputValue := range input {
		for index, char := range inputValue {
			charInt, _ := strconv.Atoi(string(char))
			if charInt == 0 {
				countArray[index] -= 1
			} else {
				countArray[index] += 1
			}
		}
	}

	gammaRate := ""
	epsilonRate := ""
	for _, value := range countArray {
		if value > 0 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gammaRateInt, _ := strconv.ParseInt(gammaRate, 2, 32)
	epsilonRateInt, _ := strconv.ParseInt(epsilonRate, 2, 32)

	fmt.Println("total : ", gammaRateInt*epsilonRateInt)
}
