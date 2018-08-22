package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

// To "implement" the interface you only need to create the properties
// declared in the interface type (method getGreeting() in this case)
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

// You can remove the receiver variable if the function don't use it:
//func (spanishBot) getGreeting() string {
func (sb spanishBot) getGreeting() string {
	// VERY custom logic for generating an spanish greeting
	return "Hola!"
}

// The interface variable cannot be passed as a receiver!!
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
