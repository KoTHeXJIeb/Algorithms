/*

Shaker Sorting Algorithm (Golang)
By CatProgrammerYT

*/

package main

import (
	"fmt"
)

func main() {
	var array []int = []int{0, 3, 5, 6, 15, 125, 25, 163, 173, 7, 1}
	fmt.Println(array)
	shakerSorting(array)
	fmt.Println(array)
}

func shakerSorting(array []int) {
	var isSwapped bool
	for i := 0; i < len(array)/2; i++ {
		for j := i; j < len(array)-i-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isSwapped = true
			}
		}
		for j := len(array) - i - 2; j > i; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
}
