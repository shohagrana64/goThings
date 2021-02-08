package main

import "fmt"

func main() {
	x := sum(5, 34)
	fmt.Println(x)
}
func sum(a float64, b float64) float64 {
	return a + b
}
