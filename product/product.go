//As it happens, I misread the assignment - this was supposed to sort, I misread multiply. Ah well, for now this is a concurrent multiplier :)

package main

import (
	"encoding/json"
	"fmt"
)

func takeInput() []int {
	fmt.Println("Please enter the values you would like to multiply surrounded by [], seperated by a comma ',' and hit enter when done - any non-integer or star values will result in the program outputting 0")
	fmt.Println("For example: [2,4,55,723,1,2]")
	var inputString string
	_, err := fmt.Scanln(&inputString)
	if err != nil {
		fmt.Printf("Unable to complete with input: %v", err)
	}
	unsortedIntList := make([]int, 1)
	json.Unmarshal([]byte(inputString), &unsortedIntList)
	return unsortedIntList
}

func splitInput(input []int) [][]int {
	splitIntList := make([][]int, 4)
	if len(input) <= 4 {
		for i := 0; i <= 4; i++ {
			input = append(input, 1)
		}
	}
	for i, v := range input {
		splitIntList[i%4] = append(splitIntList[i%4], v)
	}
	return splitIntList
}

func getProduct(listToProduct []int, result chan int) {
	product := 1
	fmt.Printf("Goroutine is producting this sub array: %v\n", listToProduct)
	for _, val := range listToProduct {
		product = product * val
	}
	result <- product
	return
}

func main() {
	fmt.Println("Sanity test")

	listOfInts := splitInput(takeInput())
	ansSlice := make([]int, 4)
	//Todo: total violation of dry, but this keeps it easy for the assignment. Yes, it bothers me.
	result0 := make(chan int, 1)
	result1 := make(chan int, 1)
	result2 := make(chan int, 1)
	result3 := make(chan int, 1)

	go getProduct(listOfInts[0], result0)
	go getProduct(listOfInts[1], result1)
	go getProduct(listOfInts[2], result2)
	go getProduct(listOfInts[3], result3)

	ansSlice[0] = <-result0
	ansSlice[1] = <-result1
	ansSlice[2] = <-result2
	ansSlice[3] = <-result3

	var final int = 1
	for _, val := range ansSlice {
		final *= val
	}

	fmt.Printf("The final product is: %d\n", final)
}
