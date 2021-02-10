package main

import (
	"fmt"
	"sort"
)

type Persons struct {
	Name string
	Age  int
}

type ByName []Persons

func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Persons{
		{"Jill", 9},
		{"Jack", 10},
		{"Rana", 5},
		{"Sakib", 11},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
