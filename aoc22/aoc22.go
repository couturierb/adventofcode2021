package aoc22

import (
	"adventofcode2021/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Ex1() {
	input := utils.ReadFileSplitLineToArray("aoc22/aoc22.txt")

	m := make(map[string]string)
	for _, line := range input {
		command, x, y, z := getCoordinates(line)
		updateMap(command, m, x, y, z)
	}
	fmt.Println("total : " + strconv.Itoa(len(m)))
}

func getCoordinates(inputLine string) (command string, x []int, y []int, z []int) {
	reg := regexp.MustCompile("-?\\d+..-?\\d+")

	command = strings.Split(inputLine, " ")[0]
	coordinates := strings.Split(inputLine, " ")[1]

	coord_x := reg.FindAllString(coordinates, -1)[0]
	coord_x_min, _ := strconv.Atoi(strings.Split(coord_x, "..")[0])
	coord_x_max, _ := strconv.Atoi(strings.Split(coord_x, "..")[1])
	for i := coord_x_min; i <= coord_x_max; i++ {
		x = append(x, i)
	}

	coord_y := reg.FindAllString(coordinates, -1)[1]
	coord_y_min, _ := strconv.Atoi(strings.Split(coord_y, "..")[0])
	coord_y_max, _ := strconv.Atoi(strings.Split(coord_y, "..")[1])
	for i := coord_y_min; i <= coord_y_max; i++ {
		y = append(y, i)
	}

	coord_z := reg.FindAllString(coordinates, -1)[2]
	coord_z_min, _ := strconv.Atoi(strings.Split(coord_z, "..")[0])
	coord_z_max, _ := strconv.Atoi(strings.Split(coord_z, "..")[1])
	for i := coord_z_min; i <= coord_z_max; i++ {
		z = append(z, i)
	}

	return
}

func updateMap(command string, m map[string]string, x []int, y []int, z []int) {
	for _, x_value := range x {
		for _, y_value := range y {
			for _, z_value := range z {
				key := strconv.Itoa(x_value) + "," + strconv.Itoa(y_value) + "," + strconv.Itoa(z_value)
				if command == "on" {
					m[key] = "on"
				} else {
					delete(m, key)
				}
			}
		}
	}
}
