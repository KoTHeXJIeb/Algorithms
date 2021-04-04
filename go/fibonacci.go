/*

Fibonacci Number Algorithm (Golang)
By CatProgrammerYT

*/

package main

func main() {

	var firstNums []int = []int{0, 1}

	for i := 0; i < len(firstNums); i++ {
		// Adding previous and second number, appending them to the list
		firstNums = append(firstNums, firstNums[i]+firstNums[i+1])
		// To make it clearer to see the list, I made the list size max of 10
		if len(firstNums) > 10 {
			break
		}
	}
}
