package main

import "fmt"

//CountArrayItemsBruteForce counts number of items in the input array
func CountArrayItemsBruteForce(input []int) int {
	count := 0
	for i := range input {
		count++
		fmt.Printf("ï: %d", i)
	}

	return count
}

func main() {

}
