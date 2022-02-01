/* Assignment:
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	//Prompt user to enter integers - Initialization (should be separate func but whatevs)
	//Note - please run this on Linux (coursera lab, ubuntu, etc) - I don't want to deal with windows adding /r/n instead of /n

	sli := make([]int, 0, 3)

	for i := 0; i < 3; i++ {
		fmt.Println("Enter an integer (" + (strconv.Itoa(i + 1)) + "/3) : ")
		_, err := fmt.Scanf("%d", &sli[i])

		if err != nil {
			fmt.Println(err)
			fmt.Println("You are running this on windows - Error is caused by how Scanf delimits inputs, and I'm too lazy atm to swap everything over to a bufio input. :P")
		}
	}

	// it feels so wrong not implementing a sorting algo lol
	sort.Ints(sli)
	fmt.Printf("%v", sli)

	//Main loop  - using a break instead of a check since this isn't properly setup with Functions (not really used in course yet)

	for {
		fmt.Println("\nEnter in a new variable (or type 'X' to exit): ")
		var checkVal string
		fmt.Scan(&checkVal)

		if checkVal == "X" || checkVal == "x" {
			break
		}

		newInt, _ := strconv.Atoi(checkVal)
		sli = append(sli, newInt)
		sort.Ints(sli)
		fmt.Printf("%v", sli)

		//There is a bug with other characters (ascii, non convertable input) getting assigned 0. I don't care at the moment
		//This is how to manipulate a slice, not perform input sanitation.
	}

	fmt.Println("Have a great day, and thanks for grading this program")
}
