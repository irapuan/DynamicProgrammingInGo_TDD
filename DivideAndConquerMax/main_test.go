package main

import "testing"

func TestMaxNumber(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		expected int
	}{
		{
			desc:     "Empty input should return zero",
			input:    []int{},
			expected: 0,
		},
		{
			desc:     "Empty input should return zero",
			input:    []int{1, 8, 5, 6},
			expected: 8,
		},
		{
			desc:     "Empty input should return zero",
			input:    []int{1, 8, 5, 6, 1, 2, 1, 1, 1, 56, 78, 2, 3, 4},
			expected: 78,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := MaxNumber(tC.input)
			if res != tC.expected {
				t.Errorf("Expected: %d, but got %d", tC.expected, res)
			}
		})
	}
}
