package main

import "fmt"

func main() {
	//simple string variable
	var x string = "Hello to my GitHub looks like You are interested <3"
	fmt.Println(x)

	//another way
	var y string
	y = "Another way to declare and print"
	fmt.Println(y)

	//variables changing values
	var z string
	z = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)

	//string concatenation
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	//shorter form of variables
	shortF := "shorter form of variables"
	fmt.Println(shortF)
	shortNumber := 5
	fmt.Println(shortNumber)
}
