package main

func FibBruteForce(input int) int {
	if input == 1 || input == 2 {
		return 1
	}
	return FibBruteForce(input-1) + FibBruteForce(input-2)
}

func Fib(input int, memo *[]int) int {
	if len(*memo) == 0 {
		*memo = append(*memo, 1)
		*memo = append(*memo, 1)
	} else if len(*memo) >= input {
		return (*memo)[input-1]
	}
	if input == 1 || input == 2 {
		return 1
	}
	recurValue := Fib(input-1, memo) + Fib(input-2, memo)
	*memo = append(*memo, recurValue)
	return recurValue
}

func main() {

}
