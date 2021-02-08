package main

import "fmt"

func main() {
	//map creation:
	vertices := make(map[string]int)

	vertices["triangles"] = 2
	vertices["square"] = 3
	vertices["decagon"] = 12

	//print map
	fmt.Println(vertices)

	//print square
	fmt.Println(vertices["square"])

	//delete square
	delete(vertices, "square")
	fmt.Println(vertices)
}
