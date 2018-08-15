package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	// 1st way to declare a struct:
	// alex := person{"Alex", "Anderson"}

	// 2nd way (better, specify the variable names):
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	// 3rd way (declare with default values -> zero values):
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
