package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func one() error {
	lc := make(chan string)

	go func() {
		if err := read("./one.txt", lc); err != nil {
			fmt.Printf("file read err %v", err)
		}
	}()

	result := 0
	for line := range lc {
		var firstDigit, lastDigit int
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				firstDigit = int(line[i] - '0')
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lastDigit = int(line[i] - '0')
				break
			}
		}

		concatenatedSum, err := strconv.Atoi((strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)))
		if err != nil {
			return err
		}

		result += concatenatedSum
	}

	return nil
}
