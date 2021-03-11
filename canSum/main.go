package main

func CanSumBruteForce(target int, options []int) bool {

	if target == 0 {
		return true
	}

	if target < 0 {
		return false
	}

	for optionsIdx := 0; optionsIdx < len(options); optionsIdx++ {
		newValue := target - options[optionsIdx]
		if CanSumBruteForce(newValue, options) {
			return true
		}
	}

	return false

}

func main() {

}
