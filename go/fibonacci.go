package main

import "fmt"

func main() {

	var firstNums []int = []int{0, 1}

	for i := 0; i < len(firstNums); i++ {
		// Adding previous and second number and appending them to the list
		firstNums = append(firstNums, firstNums[i]+firstNums[i+1])
		// To make it clearer to see list, I make list size max of 10
		if len(firstNums) > 10 {
			break
		}
	}
	fmt.Println(firstNums)
}
