package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
	var results []float64
	diskriminante := math.Pow(b, 2) - (4 * a * c)
	if diskriminante > 0 {
		results = append(results, ((-b + math.Sqrt(diskriminante)) / (2 * a)))
		results = append(results, ((-b - math.Sqrt(diskriminante)) / (2 * a)))
	} else if diskriminante == 0 {
		results = append(results, (-b / (2 * a)))
	} else {
		return nil
	}
	return results
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Printf("%f\n", computeQuadraticFormula(3, 4, 1)) //[-0.333333 -1.000000]
	fmt.Printf("%f\n", computeQuadraticFormula(2, 4, 2)) //[-1.000000]
	fmt.Printf("%f\n", computeQuadraticFormula(3, 4, 2)) //[]
}
