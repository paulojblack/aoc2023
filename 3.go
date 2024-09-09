package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func three_a() error {
	lc := make(chan string)
	counter := 0
	go func() {
		if err := read("./3.txt", lc); err != nil {
			fmt.Printf("file read err %v", err)
		}
	}()

	line2d := fill2DArr(lc)

	re, _ := regexp.Compile(`[^a-zA-Z0-9.]`)
	for yIdx, line := range line2d {
		specialCharacters := re.FindAllStringIndex(line, -1)
		fmt.Println(specialCharacters)

		for xIdx, sym := range specialCharacters {
			symPos := []int{xIdx, yIdx}
			fmt.Printf("Searching neighbors of symbol %v, at position x: %v y: %v \n", string(sym[0]), xIdx, yIdx)
			findNeighboringSum(line2d, symPos)

		}

		counter++
		if counter > 1 {

			break
		}
	}

	return nil
}

func findNeighboringSum(grid []string, pos []int) int {
	moveList := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1}}

	for _, move := range moveList {
		newPos := []int{pos[0] + move[0], pos[1] + move[1]}

		// Handle horizontal edges
		if newPos[0] < 0 || newPos[0] >= len(grid[0]) {
			continue
		}
		// Handle vertical edges
		if newPos[1] < 0 || newPos[1] >= len(grid) {
			continue
		}

		neighbor := grid[newPos[0]][newPos[1]]

		// Can directly cast as rune because all input is ASCII
		// Throw it out if its not a digit
		if !unicode.IsDigit(rune(neighbor)) {
			continue
		}

		// num := string(neighbor)

		// Check left
		// for

	}

	return 0
}

func fill2DArr(c <-chan string) []string {
	result := []string{}
	for line := range c {
		result = append(result, line)
	}

	return result
}
