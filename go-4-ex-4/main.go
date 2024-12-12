package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(c float64) float64 {
	return (c*(9/5) + 32)
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(f float64) float64 {

}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	fmt.Printf("%f\n", convertCelsiusToFahrenheit(32))
	// TODO: call the function convertFahrenheitToCelsius
	fmt.Printf(convertFahrenheitToCelsius())
}
