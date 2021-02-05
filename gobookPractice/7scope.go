package main

import "fmt"

//note that x:="Hello, World" doesn't work
//global variable needs to be declared
//x: = "Hello World"
var x string = "Hello World"

func main() {
	fmt.Println(x)
	f()
}
func f() {
	fmt.Println(x)
}
