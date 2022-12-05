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

var inputFile = flag.String("file", "day04/day04_input.txt", "Relative file path input")
var inputs []assignment

type assignment struct {
	firstSectionStart  int
	firstSectionEnd    int
	secondSectionStart int
	secondSectionEnd   int
}

func main() {
	parseInputFile()

	fmt.Println("Day04 Part 1: ", solvePartOne())
	fmt.Println("Day04  Part2: ", solvePartTwo())
}

func solvePartOne() int {
	var coverages int
	for _, assignment := range inputs {
		if assignment.firstSectionStart >= assignment.secondSectionStart && assignment.firstSectionEnd <= assignment.secondSectionEnd {
			coverages++
		} else if assignment.secondSectionStart >= assignment.firstSectionStart && assignment.secondSectionEnd <= assignment.firstSectionEnd {
			coverages++
		}
	}

	return coverages
}

func solvePartTwo() int {
	var overlappedAssignments int

	for _, assignment := range inputs {
		if assignment.firstSectionStart >= assignment.secondSectionStart && assignment.firstSectionStart <= assignment.secondSectionEnd {
			overlappedAssignments++
		} else if assignment.secondSectionStart >= assignment.firstSectionStart && assignment.secondSectionStart <= assignment.firstSectionEnd {
			overlappedAssignments++
		} else if assignment.firstSectionEnd <= assignment.secondSectionEnd && assignment.firstSectionEnd >= assignment.secondSectionStart {
			overlappedAssignments++
		} else if assignment.secondSectionEnd <= assignment.firstSectionEnd && assignment.secondSectionEnd >= assignment.firstSectionStart {
			overlappedAssignments++
		}
	}
	return overlappedAssignments
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
		sectionsAssignment := strings.Split(scanner.Text(), ",")

		firstSection := strings.Split(sectionsAssignment[0], "-")
		firstSectionStart, err := strconv.Atoi(firstSection[0])
		if err != nil {
			log.Fatal(err)
		}
		firstSectionEnd, err := strconv.Atoi(firstSection[1])
		if err != nil {
			log.Fatal(err)
		}

		secondSection := strings.Split(sectionsAssignment[1], "-")
		secondSectionStart, err := strconv.Atoi(secondSection[0])
		if err != nil {
			log.Fatal(err)
		}
		secondSectionEnd, err := strconv.Atoi(secondSection[1])
		if err != nil {
			log.Fatal(err)
		}

		inputs = append(inputs, assignment{
			firstSectionStart:  firstSectionStart,
			firstSectionEnd:    firstSectionEnd,
			secondSectionStart: secondSectionStart,
			secondSectionEnd:   secondSectionEnd,
		})
	}
}
