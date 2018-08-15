package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var result string
	for _, value := range slice {
		if value%2 != 0 {
			result = "odd"
		} else {
			result = "even"
		}

		fmt.Println(value, "is", result)
	}
}
