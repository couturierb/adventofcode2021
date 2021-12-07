package aoc7

import (
	"adventofcode2021/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
)

func Ex1() {
	input := utils.ReadFileString("aoc7/aoc7.txt")
	crabs := utils.ConvertFormattedStringToIntArray(input)
	sort.Ints(crabs)
	// fmt.Println(crabs)

	// search median
	crabsLength := len(crabs)
	var targetPosition int
	if crabsLength%2 == 0 {
		targetPosition = (crabs[crabsLength/2] + crabs[crabsLength/2-1]) / 2
	} else {
		targetPosition = crabs[crabsLength/2]
	}
	fmt.Println("target position : " + strconv.Itoa(targetPosition))

	// calcul fuel
	var fuel int
	for _, crabPosition := range crabs {
		fuel += int(math.Abs(float64(crabPosition) - float64(targetPosition)))
	}
	fmt.Println("total fuel : " + strconv.Itoa(fuel))
}

func Ex2() {
	input := utils.ReadFileString("aoc7/aoc7.txt")
	crabs := utils.ConvertFormattedStringToIntArray(input)

	// search average
	var sum int
	for _, crabPosition := range crabs {
		sum += crabPosition
	}
	targetPosition := sum / len(crabs)
	fmt.Println("target position : " + strconv.Itoa(targetPosition))

	// calcul fuel
	var fuel int
	for _, crabPosition := range crabs {
		delta := int(math.Abs(float64(crabPosition) - float64(targetPosition)))
		fuel += ((delta * delta) + delta) / 2
	}
	fmt.Println("total fuel : " + strconv.Itoa(fuel))
}
