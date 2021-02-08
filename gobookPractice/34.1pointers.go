package main

import "fmt"

func zeroo(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zeroo(&x)
	fmt.Println(x) // x is 0
}
