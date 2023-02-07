// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal.
package main

import (
	"fmt"
	"math/rand"
)

// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Hi there", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func echo(word string) {
	fmt.Println(word)
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func add(one, two, three int) int {
	return one + two + three
}

// * Write a function that returns any number
func oneNumber() int {
	value1 := rand.Intn(100)
	return value1
}

// * Write a function that returns any two numbers
func twoNumbers() (int, int) {
	value1 := rand.Intn(100)
	value2 := rand.Intn(100)
	return value1, value2
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

func main() {
	greet("Lohit")
	echo("echo")
	fmt.Println(add(1, 2, 3))
	fmt.Println(oneNumber())
	fmt.Println(twoNumbers())
	fmt.Println(add(oneNumber(), oneNumber(), oneNumber()))
}
