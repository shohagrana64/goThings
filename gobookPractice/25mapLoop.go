package main

import "fmt"

func main() {
	//map creation:
	vertices := make(map[string]int)

	vertices["triangles"] = 2
	vertices["square"] = 3
	vertices["decagon"] = 12

	for key, value := range vertices {
		fmt.Println("Key:", key, "Value:", value)
	}
}
