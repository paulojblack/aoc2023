package main

import "fmt"

func two() error {
	lc := make(chan string)

	go func() {
		if err := read("./2.txt", lc); err != nil {
			fmt.Printf("file read err %v", err)
		}
	}()

	for line := range lc {
		fmt.Println(line)
	}

	return nil
}
