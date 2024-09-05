package main

import "log"

func main() {

	err := one_two()

	if err != nil {
		log.Fatalf("we got a probbem %v", err)
	}
}
