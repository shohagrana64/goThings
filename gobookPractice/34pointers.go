package main

import "fmt"

//In this program the zero function will not modify the original x variable in the main function. But what if we wanted to? One way to do this is to use a special data type known as a pointer
func zero(x int) {
	x = 0
}
func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
}
