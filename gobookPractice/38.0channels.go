package main

import (
	"fmt"
	"time"
)

func pinger1(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func ponger1(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func printer1(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)
	go pinger1(c)
	go ponger1(c)
	go printer1(c)
	var input string
	fmt.Scanln(&input)
}
