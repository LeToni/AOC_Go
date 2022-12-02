package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var inputFile = flag.String("file", "day02/day02_input.txt", "Relative file path input")
var inputs []gameRound

type gameRound struct {
	opponent string
	player   string
}

func main() {
	parseInputFile()

	partOne := solvePartOne()
	fmt.Println("Day02 Part 1: ", partOne)

	partTwo := solvePartTwo()
	fmt.Println("Day02 Part 2:", partTwo)

}
func solvePartOne() int {
	var score int

	for _, currentRound := range inputs {
		if currentRound.player == "X" {
			score += 1
			switch currentRound.opponent {
			case "A":
				score += 3
			case "B":
				score += 0
			case "C":
				score += 6
			}
		} else if currentRound.player == "Y" {
			score += 2
			switch currentRound.opponent {
			case "A":
				score += 6
			case "B":
				score += 3
			case "C":
				score += 0
			}
		} else if currentRound.player == "Z" {
			score += 3
			switch currentRound.opponent {
			case "A":
				score += 0
			case "B":
				score += 6
			case "C":
				score += 3
			}
		}
	}

	return score
}

func solvePartTwo() int {
	var score int

	for _, round := range inputs {
		if round.opponent == "A" {
			switch round.player {
			case "X":
				score = score + 0 + 3
			case "Y":
				score = score + 3 + 1
			case "Z":
				score = score + 6 + 2
			}
		} else if round.opponent == "B" {
			switch round.player {
			case "X":
				score = score + 0 + 1
			case "Y":
				score = score + 3 + 2
			case "Z":
				score = score + 6 + 3
			}
		} else if round.opponent == "C" {
			switch round.player {
			case "X":
				score = score + 0 + 2
			case "Y":
				score = score + 3 + 3
			case "Z":
				score = score + 6 + 1
			}
		}
	}
	return score
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
		rawInputRound := strings.Split(scanner.Text(), " ")
		inputRound := gameRound{opponent: rawInputRound[0], player: rawInputRound[1]}

		inputs = append(inputs, inputRound)
	}
}
