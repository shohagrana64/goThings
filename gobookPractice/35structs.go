package main

import (
	"fmt"
	"math"
)

/*
define a Circle struct:

type Circle struct {
	x float64
	y float64
	r float64
} */

//define values of same type:

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y1 - y2
	return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func main() {
	//initialization:
	//
	//1) var c Circle
	//2) c := new(Circle)
	//
	//	initialization with variables:
	//3) c := Circle{x: 0, y: 0, r: 5}
	//if order is known:
	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)
	c.x = 10
	c.y = 5
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(circleArea(c))
	fmt.Println(c)

	fmt.Println("Now rectangle:")
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
