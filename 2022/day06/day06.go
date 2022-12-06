package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var inputFile = flag.String("file", "day06/day06_input.txt", "Relative file path input")
var inputString = flag.String("input", "", "Puzzle Input")
var datastream []rune

func main() {
	parseInput()

	fmt.Println("Day06 Part1: ", solvePartOne())
	fmt.Print("Day06 Part2: ", solvePartTwo())
}

func solvePartOne() int {
	const distinctBits int = 4

	for pos, bit := range datastream {
		trackedChars := map[rune]bool{bit: true}
		for j := pos + 1; j < pos+distinctBits; j++ {
			trackedChars[datastream[j]] = true
		}
		if len(trackedChars) == distinctBits {
			return pos + distinctBits
		}
	}
	return -1
}

func solvePartTwo() int {
	const distinctBits int = 14
	for pos, bit := range datastream {
		trackedChars := map[rune]bool{bit: true}
		for j := pos + 1; j < pos+distinctBits; j++ {
			trackedChars[datastream[j]] = true
		}
		if len(trackedChars) == distinctBits {
			return pos + distinctBits
		}
	}
	return -1
}
func parseInput() {
	flag.Parse()

	if *inputString == "" {
		file, err := os.Open(*inputFile)
		defer file.Close()

		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)

		if scanner.Scan() {
			datastream = []rune(scanner.Text())
		}
		return
	}
	datastream = []rune(*inputString)

}
