package main

import (
	"os"
)

func main() {
	file, err := os.Create("testCreate.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	file.WriteString("testing WriteString to create a file in go")
}
