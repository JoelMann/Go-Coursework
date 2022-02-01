package main

// bufio is for a Reader, fmt for print functions, OS for standard input, strings for manipulation of strings.
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Allocate space for string, prompt user for input, and assign input - Since it requires spaces as an option, fmt.Scan will not work, and will need to define a reader (since otherwise, it would need a scanner for each rune or vars for ea space input)

	fmt.Println("Please enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')

	// Error catching necessary to make the compiler happy for multi-var return on ReadString
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Create new variable for removing case mismatch, and remove trailing new lines (which are different in linux and windows, hence r and n) - Leading and trailing spaces should fail
	var lowerInput = strings.ToLower(strings.TrimRight(userInput, "\r\n"))

	//Conditional test - i, a, and n checks for prefix, contains,and suffix, with printed result.
	if strings.HasPrefix(lowerInput, "i") && strings.HasSuffix(lowerInput, "n") && strings.Contains(lowerInput, "a") {
		fmt.Println("Found!")

	} else {
		fmt.Println("Not Found!")
	}

}

/* Test Cases and Setup:

The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

To run, copy into an empty folder and cd to directory

> go run .
(or go run <path to findian.go>)

Remember to go mod init if it isn't building. Thanks for reviewing
*/
