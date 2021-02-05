package main

import "fmt"

func main() {
	fmt.Print("Feet to meters conversion....... Enter feet:")
	var ft float64
	fmt.Scanf("%f", &ft)
	var x float64 = ft * 0.3048
	fmt.Println(x)
}
