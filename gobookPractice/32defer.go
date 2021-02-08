package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {
	//defer runs the method at the end of main() or wherever it is called.

	defer second()
	first()
	fmt.Println("Hi")
}
