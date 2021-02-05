package main

import "fmt"

func main() {
	fmt.Print("Enter a Temperature in Fahrenheit:")
	var input float64
	fmt.Scanf("%f", &input)

	converted := (input - 32) * 5 / 9

	fmt.Println(converted)
}
