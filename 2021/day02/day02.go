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

var inputFile = flag.String("file", "day02/day02_input.txt", "Relative file path input")
var inputs []step

type submarine struct {
	posX int
	posY int
	aim  int
}

type step struct {
	direction string
	units     int
}

func main() {
	parseInputFile()

	partOne := solvePartOne()
	fmt.Println("Day02 Part 1:", partOne.posX*partOne.posY)

	partTwo := solvePartTwo()
	fmt.Println("Day02 Part 1:", partTwo.posX*partTwo.posY)
}

func solvePartOne() *submarine {
	var submarine submarine

	for _, command := range inputs {
		switch command.direction {
		case "forward":
			submarine.posX += command.units
		case "down":
			submarine.posY += command.units
		case "up":
			submarine.posY -= command.units
		}
	}

	return &submarine
}

func solvePartTwo() *submarine {
	var submarine submarine

	for _, command := range inputs {
		switch command.direction {
		case "forward":
			submarine.posX += command.units
			submarine.posY += submarine.aim * command.units
		case "down":
			submarine.aim += command.units
		case "up":
			submarine.aim -= command.units
		}
	}

	return &submarine
}

func parseInputFile() {
	flag.Parse()

	file, err := os.Open(*inputFile)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rawInput := strings.Split(scanner.Text(), " ")
		dir := rawInput[0]
		unit, err := strconv.Atoi(rawInput[1])

		if err != nil {
			log.Fatal(err)
		}

		inputs = append(inputs, step{direction: dir, units: unit})
	}
}
