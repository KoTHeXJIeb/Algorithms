/*
	Bubble Sorting algorithm (Golang)
	By CatProgrammerYT
*/

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// Random generator will be added later
	var array [11]int = [11]int{1, 2512, 5, 12, 51, 6, 12, 6, 17, 361, 6}
	var sortedArray [len(array)]int = bubbleSorting(array)
	// We are converting int values in string values in array
	s, _ := json.Marshal(sortedArray)

	fmt.Println(strings.Trim(string(s), "[]"))
}

func bubbleSorting(array [11]int) [11]int {
	// i - next number
	for i := 0; i < len(array); i++ {
		// j - previous number
		for j := 0; j < len(array)-1; j++ {
			// if previous is greater than previous
			if array[i] > array[j] {
				// Changing next and previous numbers in array
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
