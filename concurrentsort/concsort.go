// Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func takeInput() []int {
	fmt.Println("Please enter the values you would like to sort surrounded by [], seperated by a comma ',' and hit enter when done - any non-integer will result in the program outputting 0")
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

func swapSpot(listToSort []int, leftval int, rightval int) {
	temp := listToSort[leftval]
	listToSort[leftval] = listToSort[rightval]
	listToSort[rightval] = temp
}

func partition(listToSort []int, start int, end int) int {
	pivot := listToSort[end]
	smallIndex := start - 1

	for i := start; i <= end-1; i++ {
		if listToSort[i] < pivot {
			smallIndex += 1
			swapSpot(listToSort, smallIndex, i)
		}
	}
	swapSpot(listToSort, smallIndex+1, end)
	return (smallIndex + 1)
}

func quickSort(listToSort []int, start int, end int) {
	if start < end {
		partitionIndex := partition(listToSort, start, end)
		quickSort(listToSort, start, partitionIndex-1)
		quickSort(listToSort, partitionIndex+1, end)
	}
}

func testRandom() []int {
	var returnList []int = make([]int, 0)
	for i := 0; i < 21; i++ {
		returnList = append(returnList, rand.Intn(50))
	}
	return returnList
}

func main() {
	fmt.Println("Sanity check")
	var list []int = make([]int, 20)
	list = testRandom()
	fmt.Printf("Unsorted list is: %v", list)
	quickSort(list, 0, len(list)-1)
	fmt.Printf("Sorted list is: %v", list)
}
