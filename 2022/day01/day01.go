package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var inputFile = flag.String("file", "day01/day01_test.txt", "Relative file path input")
var input []string

func main() {

	parseInputFile()

	partOne := solvePartOne()
	fmt.Println("Day01 Part 1: ", partOne)

	partTwo := solvePartTwo()
	fmt.Println("Day01 Part 2: ", partTwo)
}

func solvePartOne() int {
	var currentCalories int
	var highestCalories int

	for _, s := range input {
		if s == "" {
			if currentCalories > highestCalories {
				highestCalories = currentCalories

			}
			currentCalories = 0
			continue
		}

		calorie, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		currentCalories += calorie
	}
	if currentCalories > highestCalories {
		highestCalories = currentCalories
	}

	return highestCalories
}

func solvePartTwo() int {
	var foundCalories []int
	var currentCalories int

	for _, s := range input {
		if s == "" {
			foundCalories = append(foundCalories, currentCalories)
			currentCalories = 0
			continue
		}

		calorie, err := strconv.Atoi(s)

		if err != nil {
			log.Fatal(err)
		}

		currentCalories += calorie
	}
	foundCalories = append(foundCalories, currentCalories)

	sort.Ints(foundCalories)

	var topThreeCalories int

	for i := 1; i <= 3; i++ {
		topThreeCalories += foundCalories[len(foundCalories)-i]
	}

	return topThreeCalories
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
		input = append(input, scanner.Text())
	}
}
