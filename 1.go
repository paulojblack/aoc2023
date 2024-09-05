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

	fmt.Println(result)
	return nil
}

var digit = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func one_two() error {
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

			if unicode.IsLetter((rune(line[i]))) {
				if i+3 < len(line) {
					if _, ok := digit[line[i:i+3]]; ok {
						firstDigit = digit[line[i:i+3]]
						break
					}
				}

				if i+4 < len(line) {
					if _, ok := digit[line[i:i+4]]; ok {
						firstDigit = digit[line[i:i+4]]
						break
					}
				}

				if i+5 < len(line) {
					if _, ok := digit[line[i:i+5]]; ok {
						firstDigit = digit[line[i:i+5]]
						break
					}
				}
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lastDigit = int(line[i] - '0')
				break
			}

			if unicode.IsLetter((rune(line[i]))) {
				if i-3 >= 0 {
					if _, ok := digit[line[i-2:i+1]]; ok {
						lastDigit = digit[line[i-2:i+1]]
						break
					}
				}

				if i-4 >= 0 {
					if _, ok := digit[line[i-3:i+1]]; ok {
						lastDigit = digit[line[i-3:i+1]]
						break
					}
				}

				if i-5 >= 0 {
					fmt.Println(line[i-4 : i+1])
					if _, ok := digit[line[i-4:i+1]]; ok {
						lastDigit = digit[line[i-4:i+1]]
						break
					}
				}
			}

		}

		concatenatedSum, err := strconv.Atoi((strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)))
		if err != nil {
			return err
		}

		result += concatenatedSum
	}

	fmt.Println(result)
	return nil
}
