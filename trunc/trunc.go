package main

import "fmt"

func main() {
	//Allocate space for float
	var userInput float64
	//Assign user input to space
	fmt.Println("Please enter in a float value to be truncated: ")
	fmt.Scan(&userInput)
	//Truncate the float - either by type conversion or string/cutoff/convert to int
	var convertedInput = int(userInput)
	fmt.Printf("Truncated value is %d", convertedInput)
}
