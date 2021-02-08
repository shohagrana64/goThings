//Write a function with one variadic parameter that finds the greatest number in a list of numbers.
package main

import "fmt"

func largestFinder(args ...int) int {
	largest := -99999
	for _, v := range args {
		if v > largest {
			largest = v
		}
	}
	return largest
}
func main() {
	xs := []int{1, 2, 3, -1, -5, 30, 5, -2}
	fmt.Println(largestFinder(xs...))
}
