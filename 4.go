package main

import "fmt"

func four_a() {
	var lc chan string
	go func() {
		if err := read("./3.txt", lc); err != nil {
			fmt.Printf("file read err %v", err)
		}
	}()

}
