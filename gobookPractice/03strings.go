package main

import "fmt"

func main() {
	fmt.Println(len("Hello World"))       //prints the length of the string 11
	fmt.Println("Hello World"[1])         //print ascii number 101
	fmt.Println(string("Hello World"[1])) //converted to string (e)
	fmt.Println("Hello " + "World")       //string concatenation
}
