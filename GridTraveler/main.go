package main

import "strconv"

func GridTravelerBruteForce(m int, n int) int {

	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	return GridTravelerBruteForce(m-1, n) + GridTravelerBruteForce(m, n-1)
}

func GridTraveler(m int, n int, memo *map[string]int) int {

	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	km := strconv.Itoa(m)
	kn := strconv.Itoa(n)
	key := km + "," + kn

	if val, ok := (*memo)[key]; ok {
		return val
	}
	res := GridTraveler(m-1, n, memo) + GridTraveler(m, n-1, memo)
	(*memo)[key] = res
	return res
}

func main() {

}
