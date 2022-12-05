package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var inputFile = flag.String("file", "day03/day03_input.txt", "Relative file path input")
var inputs []string

func main() {
	parseInputFile()

	partOne := solvePartOne()
	fmt.Println("Day03 Part 1: ", partOne)

	partTwo := solvePartTwo()
	fmt.Println("Day03 Part 2:", partTwo)
}

func solvePartOne() int {
	var priorities int

	for _, rucksack := range inputs {
		inFirstCompartment := make(map[rune]bool)
		firstCompartment := rucksack[0 : len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]

		for _, item := range firstCompartment {
			inFirstCompartment[item] = true
		}

		for _, item := range secondCompartment {
			if inFirstCompartment[item] {
				if item >= 'a' && item <= 'z' {
					priorities += int(item-'a') + 1
				} else {
					priorities += int(item-'A') + 27
				}
				break
			}
		}
	}

	return priorities
}

func solvePartTwo() int {
	var priorities int

	for i := 0; i < len(inputs); i += 3 {
		foundItems := make(map[rune]int)
		for j := 0; j < 3; j++ {
			itemTracker := make(map[rune]bool)
			for _, item := range inputs[i+j] {
				if itemTracker[item] {
					continue
				}
				itemTracker[item] = true
				foundItems[item]++
			}
		}
		for item, count := range foundItems {
			if count == 3 {
				if item >= 'a' && item <= 'z' {
					priorities += int(item-'a') + 1
				} else {
					priorities += int(item-'A') + 27
				}
				break
			}
		}
	}

	return priorities
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
		inputs = append(inputs, scanner.Text())
	}
}
