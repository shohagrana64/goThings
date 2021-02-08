package main

import "fmt"

func main() {
	//slice creation:
	x := []float64{5, 3, 4, 3, 1}
	//append function of slice
	x = append(x, 20)
	fmt.Println(x)

}
