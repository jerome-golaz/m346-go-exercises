package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Printf("%f\n", computeHypotenuse(5, 9))  //10.295630
	fmt.Printf("%f\n", computeHypotenuse(7, 29)) //29.832868
	fmt.Printf("%f\n", computeHypotenuse(13, 8)) //15.264338
}
