package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputFile = flag.String("file", "day08/day08_input.txt", "Relative file path input")
var forrest grid

type tree struct {
	height  int
	visible bool
}

type grid [][]tree

func (g grid) highestTree(x int, y int, highest *int) bool {
	height := g[y][x].height
	if height > *highest {
		*highest = height
		return true
	}

	return false
}

func main() {
	parseInput()
	fmt.Println("Day08 Part1: ", solvePartOne())
	fmt.Println("Day08 Part2: ", solvePartTwo())
}

func solvePartOne() int {
	var totalVisibleTrees int

	for y := 0; y < len(forrest); y++ {
		highestTree := 0
		for x := 0; x < len(forrest); x++ {
			if forrest.highestTree(x, y, &highestTree) {
				forrest[y][x].visible = true
			}
		}

		highestTree = 0
		for x := len(forrest) - 1; x >= 0; x-- {
			if forrest.highestTree(x, y, &highestTree) {
				forrest[y][x].visible = true
			}
		}
	}

	for x := 0; x < len(forrest); x++ {
		highestTree := 0
		for y := 0; y < len(forrest); y++ {
			if forrest.highestTree(x, y, &highestTree) {
				forrest[y][x].visible = true
			}
		}

		highestTree = 0
		for y := len(forrest) - 1; y > 0; y-- {
			if forrest.highestTree(x, y, &highestTree) {
				forrest[y][x].visible = true
			}
		}
	}

	for _, row := range forrest {
		for _, tree := range row {
			if tree.visible {
				totalVisibleTrees++
			}
		}
	}

	return totalVisibleTrees
}

func solvePartTwo() int {
	var highestScenicScore int

	for row := 0; row < len(forrest); row++ {
		for col := 0; col < len(forrest); col++ {
			givenTreeHeight := forrest[row][col].height

			highestTree := 0
			var right, left, up, down int

			for x := col + 1; x < len(forrest); x++ {
				forrest.highestTree(x, row, &highestTree)
				if highestTree >= givenTreeHeight || x == len(forrest)-1 {
					right = x - col
					break
				}
			}

			highestTree = 0
			for x := col - 1; x >= 0; x-- {
				forrest.highestTree(x, row, &highestTree)
				if highestTree >= givenTreeHeight || x == 0 {
					left = col - x
					break
				}
			}

			highestTree = 0
			for y := row - 1; y >= 0; y-- {
				forrest.highestTree(col, y, &highestTree)
				if highestTree >= givenTreeHeight || y == 0 {
					up = row - y
					break
				}
			}

			highestTree = 0
			for y := row + 1; y < len(forrest); y++ {
				forrest.highestTree(col, y, &highestTree)
				if highestTree >= givenTreeHeight || y == len(forrest)-1 {
					down = y - row
					break
				}
			}

			scenicScore := right * left * up * down

			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	return highestScenicScore
}
func parseInput() {
	flag.Parse()

	file, err := os.Open(*inputFile)

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var row []tree

		for _, i := range scanner.Text() {
			t, err := strconv.Atoi(string(i))
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, tree{
				height:  t,
				visible: false,
			})

		}
		forrest = append(forrest, row)
	}
}
