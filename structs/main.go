package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// 1st way to declare a struct:
	//alex := person{"Alex", "Anderson"}

	// 2nd way (better, specify the variable names):
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	// 3rd way (declare with default values -> zero values):
	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.print()
	(&jim).updateName("Jimmy")
	// The Go compiler performs an implicit &p on the variable,
	// so this works too (page 158 of the book)
	//jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
	// The Go compiler performs an implicit *p on the variable,
	// so this works too
	//p.firstName = newFirstName
}
