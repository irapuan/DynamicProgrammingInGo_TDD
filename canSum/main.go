package main

import (
	"strconv"
)

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

func makeKey(target int, options []int) string {
	optionsText := "|"
	for _, value := range options {
		text := strconv.Itoa(value)
		optionsText = optionsText + text + "|"
	}
	return strconv.Itoa(target) + optionsText
}

func CanSum(target int, options []int) bool {
	memory := make(map[string]bool)
	return canSum(target, options, &memory)
}

func canSum(target int, options []int, memory *map[string]bool) bool {

	if target == 0 {
		return true
	}

	if target < 0 {
		return false
	}

	key := makeKey(target, options)

	if _, ok := (*memory)[key]; ok {
		return false
	}

	for optionsIdx := 0; optionsIdx < len(options); optionsIdx++ {
		newValue := target - options[optionsIdx]

		if canSum(newValue, options, memory) {
			return true
		}
	}

	(*memory)[key] = false

	return false

}

func main() {

}
