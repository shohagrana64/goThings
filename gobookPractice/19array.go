package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	//another way to initialize array
	y := [5]float64{98, 93, 77, 82, 83}
	fmt.Println(y)

	//another way to initialize array in multiple lines
	z := [5]float64{
		98,
		93,
		77,
		82,
		83, //extra comma
	}
	fmt.Println(z)
}
