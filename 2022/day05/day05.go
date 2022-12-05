package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var inputFile = flag.String("file", "day05/day05_input.txt", "Relative file path input")
var problem = flag.String("part", "2", "Part 1/2 to solve")

var stacks = make([]stack, 9)
var procedure []rearrangement

type stack struct {
	elements []rune
}

func (s *stack) push(r []rune) {
	if *problem == "2" {
		for i := len(r); i > 0; i-- {
			s.elements = append(s.elements, r[i-1])
		}
		return
	}
	s.elements = append(s.elements, r...)
}

func (s *stack) pushToBottom(r rune) {
	s.elements = append([]rune{r}, s.elements...)
}

func (s *stack) pop(n int) []rune {

	r := s.elements[len(s.elements)-n:]
	s.elements = s.elements[:len(s.elements)-n]

	return r
}

func (s *stack) print() {
	print(string(s.elements[len(s.elements)-1]))
}

type rearrangement struct {
	amount int
	from   int
	to     int
}

func main() {
	parseInputFile()
	solveProblem()
}

func solveProblem() {

	for _, rearrangement := range procedure {
		r := stacks[rearrangement.from].pop(rearrangement.amount)
		stacks[rearrangement.to].push(r)
	}

	fmt.Print("Day05 Part ", *problem, ": ")
	for _, s := range stacks {
		if s.elements == nil || len(s.elements) == 0 {
			fmt.Print(" ")
			continue
		}
		s.print()
	}
}

func parseInputFile() {
	flag.Parse()

	file, err := os.Open(*inputFile)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan(); !strings.Contains(scanner.Text(), " 1   2   3"); scanner.Scan() {
		for i, r := range scanner.Text() {
			if r != ' ' && r != '[' && r != ']' {
				stacks[i/4].pushToBottom(r)
			}
		}
	}

	for scanner.Scan() {
		var (
			amount, from, to int
		)
		if n, _ := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &amount, &from, &to); n == 3 {
			procedure = append(procedure, rearrangement{
				amount: amount,
				from:   from - 1,
				to:     to - 1,
			})
		}
	}

}
