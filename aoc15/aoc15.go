package aoc15

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
)

var max_x int
var max_y int

type Node struct {
	visited bool
	cost    int
	weight  int
	x       int
	y       int
}

func (n Node) isEndNode() bool {
	return n.x == max_x-1 && n.y == max_y-1
}

func Ex1() {
	input := utils.ReadFileSplitLineToArray("aoc15/aoc15.txt")
	max_y = len(input)

	var grid [][]Node
	for y, value := range input {
		max_x = len(value)
		var line []Node
		for x, cell := range utils.ConvertStringToIntArray(value) {
			line = append(line, Node{cost: cell, x: x, y: y})
		}
		grid = append(grid, line)
	}

	var ready_nodes []Node
	current_node := grid[0][0]
	for !current_node.isEndNode() {
		printGrid(grid)
		new_ready_nodes := parcours(grid, current_node)
		ready_nodes = append(ready_nodes, new_ready_nodes...)
		current_node = getCurrentNode(ready_nodes)
	}

	fmt.Println("total : " + strconv.Itoa(current_node.cost))
}

func printGrid(grid [][]Node) {
	for _, line := range grid {
		for _, cell := range line {
			fmt.Print(strconv.Itoa(cell.weight) + "  ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func getCurrentNode(ready_nodes []Node) Node {
	var new_node Node
	for _, node := range ready_nodes {
		if !node.visited && (new_node.cost == 0 || new_node.cost > node.cost) {
			new_node = node
		}
	}
	return new_node
}

func parcours(grid [][]Node, current_node Node) []Node {
	var new_ready_nodes []Node
	x := current_node.x
	y := current_node.y
	grid[y][x].visited = true

	// gauche
	if x > 0 && !grid[y][x-1].visited {
		calculSibling(&grid[y][x-1], current_node, &new_ready_nodes)
	}

	// droite
	if x < len(grid[0])-1 && !grid[y][x+1].visited {
		calculSibling(&grid[y][x+1], current_node, &new_ready_nodes)
	}

	// haut
	if y > 0 && !grid[y-1][x].visited {
		calculSibling(&grid[y-1][x], current_node, &new_ready_nodes)
	}

	// bas
	if y < len(grid)-1 && !grid[y+1][x].visited {
		calculSibling(&grid[y+1][x], current_node, &new_ready_nodes)
	}

	return new_ready_nodes
}

func calculSibling(sibling *Node, current_node Node, new_ready_nodes *[]Node) {
	new_weight := sibling.cost + current_node.weight
	if sibling.weight == 0 {
		(*sibling).weight = new_weight
		*new_ready_nodes = append(*new_ready_nodes, *sibling)
	}
	if sibling.weight > new_weight {
		(*sibling).weight = new_weight
	}
}
