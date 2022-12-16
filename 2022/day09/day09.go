package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var inputFile = flag.String("file", "day09/day09_input.txt", "Relative file path input")

type position struct {
	x int
	y int
}

func (p *position) follow(head position) {
	dX := head.x - p.x
	dY := head.y - p.y
	const plankLengths = 1

	if dX <= plankLengths && dX >= -plankLengths && dY <= plankLengths && dY >= -plankLengths {
		return
	}

	if dX > 0 {
		p.x += (dX + 1) / 2
	} else {
		p.x += (dX - 1) / 2
	}

	if dY > 0 {
		p.y += (dY + 1) / 2
	} else {
		p.y += (dY - 1) / 2
	}
}

func (p *position) move(direction string, unit int) {
	switch direction {
	case "R":
		p.x += unit
	case "L":
		p.x -= unit
	case "U":
		p.y += unit
	case "D":
		p.y -= unit
	}
}

type motion struct {
	direction string
	unit      int
}

var motions []motion

func main() {
	parseInput()

	fmt.Println("Day09 Part1: ", solvePartOne())
	fmt.Println("Day10 Part2: ", solvePartTwo())
}

func solvePartOne() int {
	visited := make(map[position]bool)

	var head position
	var tail position

	visited[tail] = true
	for _, motion := range motions {

		for i := 0; i < motion.unit; i++ {
			head.move(motion.direction, 1)
			tail.follow(head)
			visited[tail] = true
		}

	}

	return len(visited)
}

func solvePartTwo() int {
	visited := make(map[position]bool)
	var knots [10]position

	for _, motion := range motions {

		for i := 0; i < motion.unit; i++ {
			knots[0].move(motion.direction, 1)

			for j := 1; j < len(knots); j++ {
				knots[j].follow(knots[j-1])
			}

			visited[knots[len(knots)-1]] = true
		}
	}
	return len(visited)
}

func parseInput() {
	flag.Parse()

	file, err := os.Open(*inputFile)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		input := strings.Split(scanner.Text(), " ")
		inputDirection := input[0]
		inputUnit, err := strconv.Atoi(input[1])

		if err != nil {
			log.Fatal(err)
		}

		motions = append(motions, motion{direction: inputDirection, unit: inputUnit})
	}
}
