package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s[0])
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
}
