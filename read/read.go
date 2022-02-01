// Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

// Submit your source code for the program, “read.go”.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//I'm not enforcing a [20]rune here because I'd rather sanitize the input than convert the type. I sanitize from the runes prior to assigning to the struct
type name struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Enter name of file to use - if not in same folder, use absolute path: ")
	var user_file string
	fmt.Scanln(&user_file)

	sli := make([]name, 0, 20)  //20 seems a reasonable start for this assignment

	f, err := os.Open(user_file)
	if err != nil {
		fmt.Println(err)
	}


	//A scanner makes more sense here than loading the entire file and parsing it. os.Read() could be used, but I didn't want to read to the space and then import past it, this just pulls up to the new line.
	scanner := bufio.NewScanner(f)

	//Will line by line for entire file. 
	for scanner.Scan() {
		raw := scanner.Text()
		var newSli name
		split_words := strings.Fields(raw)

		//This is mostly to prove I could assign the enforcement to the struct and then convert the output to string, but it feels cleaner to reference them from the struct already as strings. I could assign [20]rune to the struct from here and convert it when pushing output. This felt better.
		intermediate_fname := []rune(split_words[0])
		intermediate_lname := []rune(split_words[1])
		for len(intermediate_fname) > 20 {
			intermediate_fname = intermediate_fname[:len(intermediate_fname)-1]
		}

		for len(intermediate_lname) > 20 {
			intermediate_lname = intermediate_lname[:len(intermediate_lname)-1]
		}

		//Insertion as strings (max 20 rune) to the struct.
		newSli.fname = string(intermediate_fname)
		newSli.lname = string(intermediate_lname)

		sli = append(sli, newSli) // note append will generate new arr if allocation is gone.
	}
	
	for _, element := range sli {
    	fmt.Printf("%s", element.fname + " " + element.lname + "\n")
	}

	f.Close() // Cleaner exit case.
}
