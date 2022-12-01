package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputFile = flag.String("file", "day01/day01_input.txt", "Relative file path input")
var inputs []int

func main() {
	parseInputFile()

	partOne := solvePartOne()
	fmt.Println("Day01 Part 1: ", partOne)

	partTwo := solvePartTwo()
	fmt.Println("Day01 Part 2: ", partTwo)

}

func solvePartOne() int {
	var increaseCounter int
	previousMeasurement := inputs[0]

	for _, currentMeasurement := range inputs[1:] {
		if currentMeasurement > previousMeasurement {
			increaseCounter++
		}
		previousMeasurement = currentMeasurement
	}

	return increaseCounter
}

func solvePartTwo() int {
	lastSum := inputs[0] + inputs[1] + inputs[2]
	lastDepths := []int{inputs[0], inputs[1], inputs[2]}
	var increaseDepthCounter int

	for _, depth := range inputs[3:] {
		sum := lastSum - lastDepths[0] + depth
		lastDepths = append(lastDepths[1:], depth)

		if sum > lastSum {
			increaseDepthCounter++
		}

		lastSum = sum
	}

	return increaseDepthCounter
}

func parseInputFile() {
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		inputs = append(inputs, input)
	}
}
