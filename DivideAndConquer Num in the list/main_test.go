package main

import "testing"

func TestBruteForce(t *testing.T) {
	testCases := []struct {
		desc      string
		inputlist []int
		expected  int
	}{
		{
			desc:      "Simplest case - empty array",
			inputlist: []int{},
			expected:  0,
		},
		{
			desc:      "One item",
			inputlist: []int{6},
			expected:  1,
		},
		{
			desc:      "Three items",
			inputlist: []int{1, 2, 3},
			expected:  3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := CountArrayItemsBruteForce(tC.inputlist)
			if res != tC.expected {
				t.Errorf("Expected: %d, but got %d", tC.expected, res)
			}
		})
	}
}

func TestRecursion(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		expected int
	}{
		{
			desc:     "Simplest case",
			input:    []int{},
			expected: 0,
		},
		{
			desc:     "Simplest case",
			input:    []int{10, 23},
			expected: 2,
		},
		{
			desc:     "Simplest case",
			input:    []int{10, 23, 12, 3, 4, 5, 5, 5},
			expected: 8,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := CountArrayItems(tC.input)
			if res != tC.expected {
				t.Errorf("Expected %d , but got %d", tC.expected, res)
			}
		})
	}
}
