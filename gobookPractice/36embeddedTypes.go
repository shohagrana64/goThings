package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

//example of embedded types
type Android struct {
	//here we could use;
	//Person Person
	//but we would rather say an Android is a Person, rather than an Android has a Person.
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Name = "Rana"
	//both works
	a.Person.Talk()
	a.Talk()

}
