package main

import "fmt"

func main() {
	// Creating a variable and giving it a name
	var myName = "Lohit"
	fmt.Println("My name is ", myName)
	// Creating a variable with type notation
	var name string = "Lohit"
	fmt.Println("name =", name)
	// Creating and Assigning Value to lable shortcut
	userName := "admin"
	fmt.Println("username =", userName)
	// Creating variable without assinging value
	var sum int
	fmt.Println("The sum is ", sum)
	// Bulk creating variable and assining value
	part1, another := 1, 5
	fmt.Println("Part1 = ", part1, "Another =", another)
	// Reusing variables using shortcut
	part2, another := 2, 0
	fmt.Println("Part2 =", part2, "Another =", another)
	// Using addition
	sum = part1 + part2
	fmt.Println("Sum = ", sum)
	// Creating variable and assining values in a block
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)
	// Using block variable creation while ignoring a variable
	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}
