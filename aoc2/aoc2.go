package aoc2

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func Ex1() {
	commands := utils.ReadFileSplitToArray("aoc2/aoc2.txt")
	distance := 0
	depth := 0

	for _, command := range commands {
		commandAndValue := strings.Split(command, " ")
		value, _ := strconv.Atoi(commandAndValue[1])
		switch commandAndValue[0] {
		case "forward":
			distance += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	fmt.Println("total ex1 : ", distance*depth)
}

func Ex2() {
	commands := utils.ReadFileSplitToArray("aoc2/aoc2.txt")
	distance := 0
	depth := 0
	aim := 0

	for _, command := range commands {
		commandAndValue := strings.Split(command, " ")
		value, _ := strconv.Atoi(commandAndValue[1])
		switch commandAndValue[0] {
		case "forward":
			distance += value
			depth += value * aim
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	fmt.Println("total ex1 : ", distance*depth)
}
