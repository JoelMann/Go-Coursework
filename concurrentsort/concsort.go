// Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func takeInput() []int {
	fmt.Println("Please enter the values you would like to sort surrounded by [], seperated by a comma ',' and hit enter when done - any non-integer will result in the program outputting 0")
	fmt.Println("For example: [2,4,55,723,1,2]")
	var inputString string
	_, err := fmt.Scanln(&inputString)
	if err != nil {
		fmt.Printf("ERROR IN INPUT - Unable to complete with input: %v\n", err)
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

func asyncQuickSort(wg *sync.WaitGroup, listToSort []int, start int, end int) {
	fmt.Printf("This goroutine sorts: %v\n", listToSort)
	quickSort(listToSort, start, end)
	wg.Done()
}

func mergeTwoDimSlice(twoDimSlice [][]int) []int {
	var returnSlice = make([]int, 0)
	for i := 0; i < len(twoDimSlice); i++ {
		returnSlice = append(returnSlice, twoDimSlice[i]...)
	}
	return returnSlice
}

func main() {
	var wg sync.WaitGroup

	inputList := splitInput(takeInput())
	wg.Add(4) //note: This violates dry, but makes it easy to mark vs. making structs for it.
	go asyncQuickSort(&wg, inputList[0], 0, len(inputList[0])-1)
	go asyncQuickSort(&wg, inputList[1], 0, len(inputList[1])-1)
	go asyncQuickSort(&wg, inputList[2], 0, len(inputList[2])-1)
	go asyncQuickSort(&wg, inputList[3], 0, len(inputList[3])-1)
	wg.Wait()
	fmt.Println("Subsorting Done. Merging to slingle slice")
	finalList := mergeTwoDimSlice(inputList)
	quickSort(finalList, 0, len(finalList)-1)
	fmt.Printf("Final sorted list: %v\n", finalList)
}

//Too lazy to input a list? Here's 100 random between 1-100:
//[29,98,62,66,98,85,8,1,98,56,5,23,65,65,98,25,27,50,14,52,21,21,68,87,92,31,93,78,25,22,6,47,100,0,52,29,2,49,8,98,52,51,58,54,54,20,92,48,36,43,60,39,30,13,52,13,10,49,98,10,48,64,15,99,14,65,50,70,55,76,26,91,46,15,25,36,85,99,65,66,81,13,43,83,61,38,56,32,16,13,72,31,41,59,79,65,37,52,33,66]

//Results on that list:
/*
This goroutine sorts: [66 1 23 25 52 87 78 47 29 98 54 48 39 13 10 99 70 91 36 66 83 32 31 65 66]
This goroutine sorts: [29 98 98 65 27 21 92 25 100 2 52 54 36 30 10 48 14 55 46 85 81 61 16 41 37]
This goroutine sorts: [98 85 56 65 50 21 31 22 0 49 51 20 43 13 49 64 65 76 15 99 13 38 13 59 52]
This goroutine sorts: [62 8 5 98 14 68 93 6 52 8 58 92 60 52 98 15 50 26 25 65 43 56 72 79 33]
Subsorting Done. Merging to slingle slice
Final sorted list: [0 1 2 5 6 8 8 10 10 13 13 13 13 14 14 15 15 16 20 21 21 22 23 25 25 25 26 27 29 29 30 31 31 32 33 36 36 37 38 39 41 43 43 46 47 48 48 49 49 50 50 51 52 52 52 52 52 54 54 55 56 56 58 59 60 61 62 64 65 65 65 65 65 66 66 66 68 70 72 76 78 79 81 83 85 85 87 91 92 92 93 98 98 98 98 98 98 99 99 100]
*/
