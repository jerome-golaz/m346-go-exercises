package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints float32, maxPoints float32) float32 {
	return (gotPoints/maxPoints)*5 + 1
}

func main() {
	// TODO: call the function computeGrade
	fmt.Printf("%f\n", computeGrade(17.5, 28))  //4.125
	fmt.Printf("%f\n", computeGrade(25, 32))    //4.90625
	fmt.Printf("%f\n", computeGrade(11.25, 15)) //4.75
}
