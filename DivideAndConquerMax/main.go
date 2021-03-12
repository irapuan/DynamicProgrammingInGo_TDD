package main

//MaxNumber returns the max number in the input array in a recursive way for learning purposes
func MaxNumber(input []int) int {
	if len(input) == 0 {
		return 0
	}
	recurrentValue := MaxNumber(input[1:])
	if recurrentValue < input[0] {
		return input[0]
	}
	return recurrentValue

}

func main() {

}
