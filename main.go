package main

import "log"

func main() {

	err := two_b()

	if err != nil {
		log.Fatalf("we got a probbem %v", err)
	}
}
