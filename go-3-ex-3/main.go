package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	// TODO: Implement FizzBuzz using a for loop from Lower to Upper.
	for i := Lower; i <= Upper; i++ {
		var sAusgabe string
		if i%3 != 0 && i%5 != 0 {
			sAusgabe = string(i)
		} else {
			if i%3 == 0 {
				sAusgabe = sAusgabe + "Fizz"
			}
			if i%5 == 0 {
				sAusgabe = sAusgabe + "Buzz"
			}
		}
		fmt.Println(sAusgabe)
	}
}
