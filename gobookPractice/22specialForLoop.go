package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	//this for loop wont work because go compiler doesn't allow unused variables

	/*
		for i, value := range x {
			total += value
		}
	*/
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}