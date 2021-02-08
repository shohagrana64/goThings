package main

import "fmt"

func main() {
	//Easier better way for map creation (Without make)
	elements := map[string]string{
		"H":  "Hydrogen", //all commas
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon", //last also comma
	}

	//check if the element exists
	name, ok := elements["Un"]
	fmt.Println(name, ok)

	//if the element exists, then perform the print statement
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}
