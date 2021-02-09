package main

import (
	"fmt"
)

func hello2(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}
func main() {
	done := make(chan bool)
	go hello2(done)
	<-done //This line of code is blocking which means that until some Goroutine writes data to the done channel, the control will not move to the next line of code
	//The line of code <-done receives data from the done channel but does not use or store that data in any variable. This is perfectly legal.

	fmt.Println("main function")
}
