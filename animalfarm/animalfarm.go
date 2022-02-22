package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// This is the main "Assignment Code" - Here we are defining a type Animal, which is a struct, and declaring 3 Animals
// We also define the functions with reference pointers to the Animal type, which print the result when called.

type Animal struct {
	food         string
	locomotion   string
	spoken_sound string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Animal) Speak() {
	fmt.Println(animal.spoken_sound)
}

var cow Animal
var bird Animal
var snake Animal

// The below code is helper functions - Sanitizing user input, a help function,
// and abstracting the switch cases away from the main function.

func showUserHelp() {
	fmt.Println("This program takes 2 arguments: An animal, and a description of some part of its character or enviornment")
	fmt.Println("For example, you can check what a 'bird' would 'eat' by entering the following search term")
	fmt.Println(">bird eat  // returns 'worm'")
	fmt.Println("\nAnimals available: cow, bird, snake. Descriptions available: eat, move, speak. Use help to show this prompt. Use exit or Ctrl-C to end the program")
}

func sanitizeUserInput(input string) (string, string) {
	splitStrings := strings.Fields(input)
	if len(splitStrings) < 2 {
		splitStrings = append(splitStrings, "help", "help") //prevents index array issues. Not worrying about size at this time.
	}
	return splitStrings[0], splitStrings[1]
}

func callAnimalMethod(animal string, descriptor string) {
	// If this was any larger, I'd use a lookup table (array or map of functions) instead of a switch/if block
	// Don't say DRY - I know. But whatevs.

	switch animal {
	case "cow":
		animalMethod := cow
		if descriptor == "eat" {
			animalMethod.Eat()
		} else if descriptor == "move" {
			animalMethod.Move()
		} else if descriptor == "speak" {
			animalMethod.Speak()
		} else {
			fmt.Printf("Not a valid option for %s - try 'eat, move, or speak'\n", animal)
		}
	case "bird":
		animalMethod := bird
		if descriptor == "eat" {
			animalMethod.Eat()
		} else if descriptor == "move" {
			animalMethod.Move()
		} else if descriptor == "speak" {
			animalMethod.Speak()
		} else {
			fmt.Printf("Not a valid option for %s - try 'eat, move, or speak'\n", animal)
		}
	case "snake":
		animalMethod := snake
		if descriptor == "eat" {
			animalMethod.Eat()
		} else if descriptor == "move" {
			animalMethod.Move()
		} else if descriptor == "speak" {
			animalMethod.Speak()
		} else {
			fmt.Printf("Not a valid option for %s - try 'eat, move, or speak'\n", animal)
		}
	default:
		fmt.Println("Not a valid animal - try 'cow, bird, or snake'")
	}
}

// The main function inits the 3 cases in the assingment (cow bird snake), prompts for input, and runs the event loop.
// The loop is infinite, but I've manually put in an exit clause for sanity.

func main() {

	//Init in main for global scope the 3 animals. Note this isn't fully dynamic - I'm using switch cases to handle user input.
	cow = Animal{food: "grass", locomotion: "walk", spoken_sound: "moo"}
	bird = Animal{food: "worm", locomotion: "fly", spoken_sound: "peep"}
	snake = Animal{food: "mice", locomotion: "slither", spoken_sound: "hsss"}

	//Intro text to program
	fmt.Println("Enter in promts as '>animal description' separated by a space, then hit enter. For help, type '>help'")
	fmt.Println("To exit, hit Ctrl-C or use '>exit'")

	exitCondition := false // changes to 'true' with 'exit' input

	// Infinite loop starts here
	for !exitCondition {
		fmt.Printf(">")

		//Using a scanner to allocate from StdIn with spaces, preventing (some) input entry bugs from Scanf
		var user_input string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		user_input = scanner.Text() // Ensures only 'string' is returned regardless of input

		input_1, input_2 := sanitizeUserInput(user_input)

		if input_1 == "help" {
			showUserHelp()
		} else if input_1 == "exit" || input_2 == "exit" {
			// Edge case if user input somehow takes first arg - second will also allow breaking the loop (Windows specific)
			exitCondition = true
		} else {
			callAnimalMethod(input_1, input_2) // this calls the relevant animal's Eat(), Move(), or Speak() method
		}
	}
}
