package main

import (
	"fmt"
	"strconv"
	"strings"
)

type cubeConfig struct {
	Color string
}

var totalCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func two_a() error {
	lc := make(chan string)

	go func() {
		if err := read("./2.txt", lc); err != nil {
			fmt.Printf("file read err %v", err)
		}
	}()

	totalSum := 0
	for line := range lc {
		countMeIn := true
		parts := strings.Split(line, ":")

		id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
		rolls := strings.Split(parts[1], ";")

		for _, roll := range rolls {
			cubeRolls := strings.Split(roll, ",")

			for _, cubeCount := range cubeRolls {
				cubeParts := strings.Split(strings.TrimLeft(cubeCount, " "), " ")
				intCubeRollCount, _ := strconv.Atoi(cubeParts[0])

				if totalCubes[cubeParts[1]] < intCubeRollCount {
					countMeIn = false
				}
			}
		}

		if countMeIn == true {
			totalSum += id
		}

	}
	fmt.Println(totalSum)
	return nil
}

func two_b() error {
	lc := make(chan string)

	go func() {
		if err := read("./2.txt", lc); err != nil {
			fmt.Printf("file read err %v", err)
		}
	}()

	totalSum := 0
	for line := range lc {

		parts := strings.Split(line, ":")

		rolls := strings.Split(parts[1], ";")

		rollMap := map[string]int{
			"red":   1,
			"blue":  1,
			"green": 1,
		}
		for _, roll := range rolls {
			cubeRolls := strings.Split(roll, ",")

			for _, cubeCount := range cubeRolls {
				cubeParts := strings.Split(strings.TrimLeft(cubeCount, " "), " ")
				intCubeRollCount, _ := strconv.Atoi(cubeParts[0])

				if rollMap[cubeParts[1]] < intCubeRollCount {
					rollMap[cubeParts[1]] = intCubeRollCount
				}

			}
		}
		multiplier := 1

		for _, v := range rollMap {
			multiplier *= v
		}
		totalSum += multiplier
	}
	fmt.Println(totalSum)
	return nil
}
