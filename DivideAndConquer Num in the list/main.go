package main

import "fmt"

//CountArrayItemsBruteForce counts number of items in the input array
func CountArrayItemsBruteForce(input []int) int {
	count := 0
	for i := range input {
		count++
		fmt.Printf("ï: %d\n", i)
	}

	return count
}

//CountArrayItems return the number of items in the input array
func CountArrayItems(input []int) int {
	if len(input) == 0 {
		return 0
	}
	return 1 + CountArrayItems(input[1:])
}

func main() {

}
