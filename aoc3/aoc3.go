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

func Ex2() {
	input := utils.ReadFileSplitToArray("aoc3/aoc3.txt")

	oxygenGenerator := findOxygen(input, 0)
	co2Scrubber := findco2Scrubber(input, 0)

	fmt.Println(oxygenGenerator)
	fmt.Println(co2Scrubber)

	oxygenGeneratorInt, _ := strconv.ParseInt(oxygenGenerator, 2, 32)
	co2ScrubberInt, _ := strconv.ParseInt(co2Scrubber, 2, 32)

	fmt.Println("total : ", oxygenGeneratorInt*co2ScrubberInt)
}

func findOxygen(input []string, charPosition int) string {
	var arrayWith1 []string
	var arrayWith0 []string

	for _, inputValue := range input {
		if inputValue[charPosition:charPosition+1] == "1" {
			arrayWith1 = append(arrayWith1, inputValue)
		} else {
			arrayWith0 = append(arrayWith0, inputValue)
		}
	}

	if charPosition >= len(input[0])-1 {
		if len(arrayWith1) > 0 {
			return arrayWith1[0]
		} else {
			return arrayWith0[0]
		}
	}

	if len(arrayWith1) >= len(arrayWith0) {
		return findOxygen(arrayWith1, charPosition+1)
	} else {
		return findOxygen(arrayWith0, charPosition+1)
	}
}

func findco2Scrubber(input []string, charPosition int) string {
	var arrayWith1 []string
	var arrayWith0 []string

	for _, inputValue := range input {
		if inputValue[charPosition:charPosition+1] == "1" {
			arrayWith1 = append(arrayWith1, inputValue)
		} else {
			arrayWith0 = append(arrayWith0, inputValue)
		}
	}

	if charPosition >= len(input[0])-1 {
		if len(arrayWith0) > 0 {
			return arrayWith0[0]
		} else {
			return arrayWith1[0]
		}
	}

	if len(arrayWith1) == 0 || (len(arrayWith0) > 0 && len(arrayWith0) <= len(arrayWith1)) {
		return findco2Scrubber(arrayWith0, charPosition+1)
	} else {
		return findco2Scrubber(arrayWith1, charPosition+1)
	}
}
