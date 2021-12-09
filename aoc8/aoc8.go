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

	var total int
	for _, inputValue := range input {
		valueRightAndLeft := strings.Split(inputValue, " | ")
		dictionary := createDictionary(valueRightAndLeft[0])
		valuesRight := strings.Split(valueRightAndLeft[1], " ")

		digit1, _ := findDigit(valuesRight[0], dictionary)
		digit2, _ := findDigit(valuesRight[1], dictionary)
		digit3, _ := findDigit(valuesRight[2], dictionary)
		digit4, _ := findDigit(valuesRight[3], dictionary)

		total += digit1*1000 + digit2*100 + digit3*10 + digit4
	}

	fmt.Println(total)
}

func createDictionary(input string) [10]string {
	dict := [10]string{}
	values := strings.Split(input, " ")

	// On parcourt une première fois pour trouver 1, 4, 7 et 8
	for _, value := range values {
		digit, err := findDigit(value)
		if err == nil {
			dict[digit] = value
		}
	}

	// On parcourt une 2ème fois pour trouver 0, 2, 3, 5, 6 et 9
	for _, value := range values {
		digit, err := findDigit(value, dict)
		if err == nil {
			dict[digit] = value
		}
	}

	return dict
}

func findDigit(value string, otherParams ...[10]string) (int, error) {
	switch len(value) {
	case 2:
		return 1, nil
	case 3:
		return 7, nil
	case 4:
		return 4, nil
	case 5:
		// 3 = 5 char et contient les segments de 1
		// 5 = 5 char et contient 3 segments de 4
		// 2 sinon
		if len(otherParams) > 0 {
			if strings.Contains(value, otherParams[0][1][0:1]) && strings.Contains(value, otherParams[0][1][1:2]) {
				return 3, nil
			} else {
				var howManySegments int
				for i := 0; i < 4; i++ {
					howManySegments += strings.Count(value, otherParams[0][4][i:i+1])
				}
				if howManySegments == 3 {
					return 5, nil
				} else {
					return 2, nil
				}
			}
		}
	case 6:
		// 9 = 6 char et contient les segments de 4
		// 0 = 6 char et contient les segments de 1 (et n'est pas 9 du coup)
		// 6 sinon
		if len(otherParams) > 0 {
			if strings.Contains(value, otherParams[0][4][0:1]) && strings.Contains(value, otherParams[0][4][1:2]) &&
				strings.Contains(value, otherParams[0][4][2:3]) && strings.Contains(value, otherParams[0][4][3:4]) {
				return 9, nil
			} else if strings.Contains(value, otherParams[0][1][0:1]) && strings.Contains(value, otherParams[0][1][1:2]) {
				return 0, nil
			} else {
				return 6, nil
			}
		}
	case 7:
		return 8, nil
	}

	return -1, errors.New("Pas de valeur trouvée")
}
