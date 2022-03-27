// Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

package main

import (
	"fmt"
)

func takeInput() []int {
	unsortedIntList := make([]int, 0)
	return unsortedIntList
}

func splitInput(input []int, numberOfSplits int) [][]int {
	splitIntList := make([][]int, numberOfSplits)
	return splitIntList
}

func getProduct(listToProduct []int) int {
	product := 1
	for _, val := range listToProduct {
		product = product * val
	}
	return product
}


func main() {
	fmt.Println("Sanity test")

	listOfInts := splitInput(takeInput(), 4)
	ansSlice := make([]int,len(listOfInts))
//Todo: possibly create a constructor for a variable amount of concurrent functions. 
	ansSlice[0] := go getProduct(listOfInts[0])
	ansSlice[1] := go getProduct(listOfInts[1])
	ansSlice[2] := go getProduct(listOfInts[2])
	ansSlice[3] := go getProduct(listOfInts[3])

	fmt.Println(getProduct(ansSlice))

}