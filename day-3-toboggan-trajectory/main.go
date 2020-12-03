package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	rightCoordinates := []int{3}
	downCoordinates := []int{1}
	fmt.Print("Part 1: ")
	countTrees(rightCoordinates, downCoordinates)

	rightCoordinates = []int{1, 3, 5, 7, 1}
	downCoordinates = []int{1, 1, 1, 1, 2}
	fmt.Print("Part 2: ")
	countTrees(rightCoordinates, downCoordinates)
}

func countTrees(rightCoordinates []int, downCoordinates []int) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	treesCounter := make([]int, len(rightCoordinates))

	rowCounter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		for i, right := range rightCoordinates {
			if rune(row[right * rowCounter / downCoordinates[i] % len(row)]) == '#' && rowCounter % downCoordinates[i] == 0 {
				treesCounter[i] += 1
			}
		}
		rowCounter += 1
	}

	product := 1
	for _, x := range treesCounter {
		product *= x
	}
	fmt.Println("Encountered trees", treesCounter, "=", product)
}
