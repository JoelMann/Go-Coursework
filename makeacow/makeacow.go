package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Animals implementing the Animal interface - Cow, Bird, Snake
type Cow struct {
	food       string
	locomotion string
	sound      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}
func (c Cow) Speak() {
	fmt.Println(c.sound)
}

type Bird struct {
	food       string
	locomotion string
	sound      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}
func (b Bird) Move() {
	fmt.Println(b.locomotion)
}
func (b Bird) Speak() {
	fmt.Println(b.sound)
}

type Snake struct {
	food       string
	locomotion string
	sound      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}
func (s Snake) Move() {
	fmt.Println(s.locomotion)
}
func (s Snake) Speak() {
	fmt.Println(s.sound)
}

// Functions implementing the logic abstraction to determine animal creation/queries
func NewAnimal(name string, animal_type string, db map[string]Animal) {
	if db[name] != nil {
		fmt.Println("This already exists - please run '>query name action' for information.")
	} else {
		switch animal_type {
		case "cow":
			db[name] = Cow{food: "grass", locomotion: "walk", sound: "moo"}
			fmt.Println("Created it!")
		case "bird":
			db[name] = Bird{food: "worms", locomotion: "fly", sound: "peep"}
			fmt.Println("Created it!")
		case "snake":
			db[name] = Snake{food: "mice", locomotion: "slither", sound: "hsss"}
			fmt.Println("Created it!")
		default:
			fmt.Println("That animal type has not yet been added - please try only 'cow, bird, snake' as options")
		}
	}
}

func Query(name string, action string, db map[string]Animal) {
	if db[name] == nil {
		fmt.Println("No animal with that name has been created. Use '>newanimal name type' if you'd like to create it.")
	} else {
		switch action {
		case "eat":
			db[name].Eat()
		case "move":
			db[name].Move()
		case "speak":
			db[name].Speak()
		default:
			fmt.Println("That action isn't added - please ensure you are only asking for the actions 'eat, move, speak' in '>query animal action'")
		}
	}
}

// Other helper functions:
func showUserHelp() {
	fmt.Println("This program takes 2 options: newanimal, or query")
	fmt.Println("Create a new animal of the types cow, bird, or snake")
	fmt.Println(">newanimal bob cow  // adds bob the cow to the db")
	fmt.Println("\nAnimals available: cow, bird, snake. Actions available to query: eat, move, speak. Use help to show this prompt. Use exit or Ctrl-C to end the program")
}

func sanitizeUserInput(input string) (string, string, string) {
	splitStrings := strings.Fields(input)
	if len(splitStrings) < 3 {
		splitStrings = append(splitStrings, "help", "help", "help") //prevents index array issues. Not worrying about size at this time.
	}
	return strings.ToLower(splitStrings[0]), strings.ToLower(splitStrings[1]), strings.ToLower(splitStrings[2])
}

// CLI logic and application
func main() {

	animal_map := make(map[string]Animal)

	//Intro text to program
	fmt.Println("Enter in prompts as '>newanimal name animal' or '>query name action' separated by a space, then hit enter. For help, type '>help'")
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

		user_command, animal_name, input_3 := sanitizeUserInput(user_input)

		switch user_command {
		case "help":
			showUserHelp()
		case "exit":
			exitCondition = true
		case "newanimal":
			NewAnimal(animal_name, input_3, animal_map)
		case "query":
			Query(animal_name, input_3, animal_map)
		default:
			fmt.Println("Please ensure you are using '>newanimal' or '>query'")
		}
	}
}
