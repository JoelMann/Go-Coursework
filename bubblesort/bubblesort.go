// Write a Bubble Sort program in Go. The program
// should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to
// greatest. Use your favorite search tool to find a description of how the bubble
// sort algorithm works.

// As part of this program, you should write a
// function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
// order.

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements in the
// slice. You should write a Swap() function which performs this operation. Your Swap()
// function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice. The Swap() function should return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.

package main

import "fmt"

func Swap(slice []int, i int) {
	var tempVal int = slice[i]
	slice[i] = slice[i+1]
	slice[i+1] = tempVal
}

func BubbleSort(numbers []int) {
	counter := 0
	for i := 0; i <= len(numbers)-2; i++ {
		if numbers[i] > numbers[i+1] {
			Swap(numbers, i)
			counter += 1
		}
	}
	if counter > 0 {
		BubbleSort(numbers)
	}

}

func main() {
	UnsortedSlice := make([]int, 0, 10)
	fmt.Println("Enter in integers one at a time, each followed by an enter. The program will return a slice after the 10th entry, or the first non-integer value/empty character entered")
	fmt.Println("For Example:\n1\n6\n-24\n ... Enter your Values ...")
	for i := 0; i < 10; i++ {
		var val int
		_, err := fmt.Scanf("%d", &val)
		if err != nil {
			break
		}
		UnsortedSlice = append(UnsortedSlice, val)
	}

	// test := []int{1,2,3,9,2,5,4,9,199,102,-1}
	fmt.Println("The values entered:")
	fmt.Printf("%v", UnsortedSlice)
	BubbleSort(UnsortedSlice)
	fmt.Println("\nSorted:")
	fmt.Printf("%v", UnsortedSlice)

}
