package aoc6

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
)

const MAX_DAYS1 = 80
const MAX_DAYS2 = 256

func Ex1() {
	input := utils.ReadFileString("aoc6/aoc6.txt")
	lanternfish := utils.ConvertFormattedStringToIntArray(input)

	for i := 0; i < MAX_DAYS1; i++ {
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
	}

	fmt.Println("total : " + strconv.Itoa(len(lanternfish)))
}

func Ex2() {
	input := utils.ReadFileString("aoc6/aoc6.txt")
	lanternfish := utils.ConvertFormattedStringToIntArray(input)
	lanternfish2 := [9]int{}

	for _, fish := range lanternfish {
		lanternfish2[fish]++
	}

	for i := 0; i < MAX_DAYS2; i++ {
		howManyNewFish := 0
		for j := 0; j < len(lanternfish2)-1; j++ {
			if j == 0 {
				howManyNewFish = lanternfish2[0]
			}
			lanternfish2[j] = lanternfish2[j+1]
		}
		lanternfish2[6] += howManyNewFish
		lanternfish2[8] = howManyNewFish
	}

	total := 0
	for _, value := range lanternfish2 {
		total += value
	}
	fmt.Println("total : " + strconv.Itoa(total))
}
