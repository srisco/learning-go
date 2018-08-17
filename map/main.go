package main

import (
	"fmt"
)

func main() {
	// Initialize an empty map (1)
	//var colors map[string]string

	// Initialize an empty map (2)
	//colors := make(map[string]string)

	// To delete a key (and value)
	//delete(colors, "key")

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
