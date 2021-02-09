package main

import "fmt"

func fo(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go fo(i)
	}
	var input string
	fmt.Scanln(&input)
}
