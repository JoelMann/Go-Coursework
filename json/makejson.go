// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

// Submit your source code for the program,
// “makejson.go”.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//Create a structure to allocate (because why not)
type Person struct {
	name    string
	address string
}

func main() {

	user := new(Person)
	//Need to allocate something better than fmt scan
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a name: ")
	user.name, _ = reader.ReadString('\n')
	user.name = strings.TrimRight(user.name, "\r\n")
	
	fmt.Println("Enter an address: ")
	user.address, _ = reader.ReadString('\n')
	user.address = strings.TrimRight(user.address, "\r\n")

	//Map the dereferenced object fields to a hash table from the user object

	mapped_person := map[string]string{"name": user.name, "address": user.address}

	output, _ := json.Marshal(mapped_person)

	fmt.Println(string(output))
	//note - I'm type casting to print as string, otherwise you will get a raw byte arr - Format is JSON tho, even as a byte array
}
