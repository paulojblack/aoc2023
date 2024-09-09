package main

import "log"

func main() {

	err := three_a()

	if err != nil {
		log.Fatalf("we got a probbem %v", err)
	}
}
