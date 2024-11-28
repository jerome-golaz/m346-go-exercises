package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	// fmt.Println("the dice shows", eyes, "eyes")
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	// fmt.Println("the dice was rolled at", when)
	fmt.Fprintln(os.Stderr, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go TODO
	// go run ex3/main.go >eyes.txt
	// go run ex3/main.go 2>dice.log
}
