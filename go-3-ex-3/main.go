package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	// TODO: Implement FizzBuzz using a for loop from Lower to Upper.
	for i := Lower; i <= Upper; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Printf("Fizz ")
		} else if i%5 == 0 {
			fmt.Printf("Buzz ")
		} else {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}
}
