package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("43test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	p := strings.Split(str, " ")
	for i := 0; i < len(p); i++ {
		fmt.Println(i, p[i])
	}

}
