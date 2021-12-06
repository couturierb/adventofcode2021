package aoc6

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
)

const MAX_DAYS = 80

func Ex1() {
	input := utils.ReadFileString("aoc6/aoc6.txt")
	lanternfish := utils.ConvertFormattedStringToIntArray(input)

	for i := 0; i < MAX_DAYS; i++ {
		howManyNewFish := 0

		for index, value := range lanternfish {
			lanternfish[index] = value - 1
			if lanternfish[index] < 0 {
				lanternfish[index] = 6
				howManyNewFish++
			}
		}

		for j := 0; j < howManyNewFish; j++ {
			lanternfish = append(lanternfish, 8)
		}

		// fmt.Println(lanternfish)
	}

	fmt.Println("total : " + strconv.Itoa(len(lanternfish)))
}
