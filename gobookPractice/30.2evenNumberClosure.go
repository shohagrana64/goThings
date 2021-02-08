package main

import "fmt"

func makeEvenGenerator() func() (ret int) {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}
func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}
