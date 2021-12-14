package aoc14

import (
	"adventofcode2021/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const STEP_NB = 40

func Ex1() {
	input := utils.ReadFileSplitLineToArray("aoc14/aoc14.txt")
	template, pairInsertion := createPairInsertion(input)

	// fmt.Println(template)
	for i := 0; i < STEP_NB; i++ {
		fmt.Println(i)
		template = step(template, pairInsertion)
		// fmt.Println(template)
	}

	// fmt.Println("most : " + strconv.Itoa(getMostCommon(template)))
	// fmt.Println("least : " + strconv.Itoa(getLeastCommon(template)))
	fmt.Println("total : " + strconv.Itoa(getMostCommon(template)-getLeastCommon(template)))
}

func getMostCommon(template string) int {
	mostCommon := 0
	for _, value := range template {
		count := strings.Count(template, string(value))
		if count > mostCommon {
			mostCommon = count
		}
	}
	return mostCommon
}

func getLeastCommon(template string) int {
	leastCommon := 999999999999999999
	for _, value := range template {
		count := strings.Count(template, string(value))
		if count < leastCommon {
			leastCommon = count
		}
	}
	return leastCommon
}

func step(template string, pairInsertion map[string]string) string {
	var newTemplate string

	for i := 0; i < len(template)-1; i++ {
		newTemplate += pairInsertion[template[i:i+2]]
	}

	newTemplate += string(template[len(template)-1])
	return newTemplate
}

func createPairInsertion(input []string) (string, map[string]string) {
	pairInsertion := make(map[string]string)
	reg := regexp.MustCompile(" -> ")
	var template string

	for index, value := range input {
		if index == 0 {
			template = value
		} else if index != 1 {
			values := reg.Split(value, -1)
			pairInsertion[values[0]] = values[0][:1] + values[1]
		}
	}

	return template, pairInsertion
}
