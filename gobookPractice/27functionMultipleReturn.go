package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := (sqrt(16))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

//function that takes 1 number and return 2 values the square root and the error if any
func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return math.Sqrt(a), nil

}
