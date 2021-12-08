package aoc8

import (
	"adventofcode2021/utils"
	"errors"
	"fmt"
	"strings"
)

func Ex1() {
	input := utils.ReadFileSplitToArray("aoc8/aoc8.txt")
	var howManyDigit int

	for _, value := range input {
		valueRight := strings.Split(value, " | ")[1]
		valuesRight := strings.Split(valueRight, " ")
		for _, valueRight := range valuesRight {
			if (len(valueRight) >= 2 && len(valueRight) <= 4) || len(valueRight) == 7 {
				howManyDigit++
			}
		}
	}

	fmt.Println(howManyDigit)
}

func Ex2() {
	input := utils.ReadFileSplitToArray("aoc8/aoc8.txt")

	for _, value := range input {
		valueRightAndLeft := strings.Split(value, " | ")
		dictionary := createDictionary(valueRightAndLeft[0])
		fmt.Println(dictionary)
		// valuesRight := strings.Split(valueRight, " ")
		// for _, valueRight := range valuesRight {
		// 	if (len(valueRight) >= 2 && len(valueRight) <= 4) || len(valueRight) == 7 {
		// 		howManyDigit++
		// 	}
		// }
	}
}

func createDictionary(input string) [9]string {
	dict := [9]string{}
	values := strings.Split(input, " ")

	// On parcourt une première fois pour trouver 1, 4, 7 et 8
	for _, value := range values {
		digit := findDigit(value)
		dict[digit] = value
	}

	// On parcourt une 2ème fois pour trouver 2, 3, 5, 6 et 9
	for _, value := range values {
		digit := findDigit(value, dict)
		dict[digit] = value
	}

	return dict
}

func findDigit(value string, dict ...[9]string) (int, error) {
	// 1 = 2 digit
	// 4 = 4 digit
	// 7 = 3 digit
	// 8 = 7 digit
	// 2, 3, 5 = 5 digits and 2 not letters
	// 6, 9 = 6 digit and 1 nat letter
	switch len(value) {
	case 2:
		return 1, nil
	case 3:
		return 7, nil
	case 4:
		return 4, nil
	case 5:
		if len(dict) > 0 {
			// TODO
		}
	case 6:
		if len(dict) > 0 {
			// TODO
		}
	case 7:
		return 8, nil
	}

	return -1, errors.New("blabla")
}
