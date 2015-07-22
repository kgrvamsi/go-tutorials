package main

import (
	"fmt"
)

// Multiple values based Functions
func addition() (int, int) {
	return 4, 8
}

//Variadic Functions

func division(messages string, values ...int) {

	total := 0
	for _, value := range values {

		total += value
	}

	fmt.Println(messages, values, total)

}

func substraction(args ...string) {
	fmt.Println(args)
}

func main() {

	a, b := addition()

	fmt.Println(a, b)

	division("The Values are :", 1, 2)
	division("The Values are :", 1, 2, 3)
	division("The Values are :", 4, 8, 16)

	substraction("Hello", "How", "Are", "You")

}
