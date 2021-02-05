package main

import "fmt"

func main() {
	//Go has only one loop =====> the for loop
	//while loop using for loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	//for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
