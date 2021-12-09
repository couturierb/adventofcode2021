package aoc4

import (
	"adventofcode2021/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	Name        string
	Cases       [5][5]int
	CurrentLine int  `0`
	win         bool `false`
}

func (b *Board) addLine(line []string) {
	for k, v := range line {
		intValue, _ := strconv.Atoi(v)
		b.Cases[b.CurrentLine][k] = intValue
	}
	b.CurrentLine++
}

func (b *Board) mark(value int) {
	for i, line := range b.Cases {
		for j, oneCase := range line {
			if oneCase == value {
				b.Cases[i][j] = -1
			}
		}
	}
}

func (b *Board) check() bool {
	if b.win {
		return false
	}

	for i := 0; i < len(b.Cases); i++ {
		checkLine := 0
		checkColumn := 0
		for j := 0; j < len(b.Cases); j++ {
			checkColumn += b.Cases[i][j]
			checkLine += b.Cases[j][i]
		}

		if checkLine == -5 || checkColumn == -5 {
			b.win = true
			return true
		}
	}

	return false
}

func (b *Board) total(winValue int) int {
	total := 0

	for _, line := range b.Cases {
		for _, oneCase := range line {
			if oneCase != -1 {
				total += oneCase
			}
		}
	}

	return total * winValue
}

func (b Board) Print() {
	fmt.Println(" --- Board " + b.Name + " ---")
	for _, line := range b.Cases {
		fmt.Println(line)
	}
}

func Ex1Ex2() {
	input := utils.ReadFileSplitLineToArray("aoc4/aoc4.txt")
	var tirage []string
	boards := []Board{}
	currentBoard := -1

	for index, line := range input {
		if index == 0 {
			// Première ligne avec le tirage
			tirage = strings.Split(line, ",")
		} else if line == "" {
			// Ligne vide, on crée un nouveau Board
			currentBoard++
			boards = append(boards, Board{Name: "board " + strconv.Itoa(currentBoard)})
		} else {
			// On ajoute les lignes au board courant
			// lineValues := strings.Split(line, " ")
			re := regexp.MustCompile("[0-9]+")
			lineValues := re.FindAllString(line, -1)
			boards[currentBoard].addLine(lineValues)
		}
	}

	for _, value := range tirage {
		for boardIndex, board := range boards {
			intValue, _ := strconv.Atoi(value)
			boards[boardIndex].mark(intValue)
			if boards[boardIndex].check() {
				total := boards[boardIndex].total(intValue)
				fmt.Println("board " + board.Name + " win with " + strconv.Itoa(total))
			}
		}
	}
}
