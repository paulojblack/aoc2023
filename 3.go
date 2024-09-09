package main

import (
	"fmt"
	"regexp"
	"strings"
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

		for _, sym := range specialCharacters {
			xIdx := sym[0]
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
		newX := newPos[0]
		newY := newPos[1]
		// Handle horizontal edges
		if newX < 0 || newX >= len(grid[0]) {
			continue
		}
		// Handle vertical edges
		if newY < 0 || newY >= len(grid) {
			continue
		}

		neighbor := grid[newY][newX]

		// Can directly cast as rune because all input is ASCII
		// Throw it out if its not a digit
		if !unicode.IsDigit(rune(neighbor)) {
			continue
		}

		num := string(neighbor)
		byteDot := "."
		byteDot = string(byteDot)
		// Check left

		for i := newX - 1; i >= 0 && unicode.IsDigit(rune(grid[newY][i])); i-- {
			newNum := string(grid[newY][i])
			grid[newY] = strings.Replace(grid[newY], newNum, ".", -1)
			num = newNum + num
		}

		fmt.Println(grid[newY])

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
