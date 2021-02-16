package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "whats up my boy"
	texts := strings.Split(text, " ")
	fmt.Println(texts)
}
