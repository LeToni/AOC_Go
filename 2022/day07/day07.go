package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var inputFile = flag.String("file", "day07/day07_input.txt", "Relative file path input")
var root dir

type dir struct {
	name     string
	size     int
	parent   *dir
	children []*dir
}

func (d dir) crawl(f func(dir)) {
	f(d)
	for _, child := range d.children {
		child.crawl(f)
	}
}

func main() {
	parseInput()
	fmt.Println("Day07 Part 1: ", solvePartOne())
	fmt.Println("Day07 Part 2: ", solvePartTwo())
}

func solvePartOne() int {
	var totalSum int

	partOneCondition := func(d dir) {
		if d.size <= 100000 {
			totalSum += d.size
		}
	}

	root.crawl(partOneCondition)

	return totalSum
}

func solvePartTwo() int {
	smallestSize := math.MaxInt
	const diskSpace int = 70000000
	const minRequiredUnusedSpace int = 30000000
	const availableSpace int = diskSpace - minRequiredUnusedSpace

	partTwoCondition := func(d dir) {
		if root.size-d.size < availableSpace && smallestSize > d.size {
			smallestSize = d.size
		}
	}

	root.crawl(partTwoCondition)
	return smallestSize
}

func parseInput() {
	flag.Parse()

	file, err := os.Open(*inputFile)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		if n, _ := fmt.Sscanf(scanner.Text(), "$ cd /"); n == 1 {
			log.Fatal("File cmd error line 1, not starting with: $ cd /")
		}
	}

	currentDir := &root
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "$" {
			if line[1] == "ls" {
				continue
			} else if line[1] == "cd" {
				if line[2] == ".." {
					currentDir = currentDir.parent
					continue
				}
				var subDir dir
				subDir.name = line[2]
				currentDir.children = append(currentDir.children, &subDir)
				subDir.parent = currentDir

				currentDir = &subDir
			}

		} else if line[0] == "dir" {
			continue
		} else {
			fileSize, err := strconv.Atoi(line[0])
			if err != nil {
				log.Fatal("Fail to parse file size")
			}

			dirToUpdate := currentDir
			for dirToUpdate != nil {
				dirToUpdate.size += fileSize
				dirToUpdate = dirToUpdate.parent
			}
		}
	}
}
