package main

import "testing"

func TestCanSumBruteForce(t *testing.T) {
	testCases := []struct {
		desc     string
		target   int
		options  []int
		expected bool
	}{
		{
			desc:     "Simplest case",
			target:   7,
			options:  []int{3, 4},
			expected: true,
		},
		{
			desc:     "Simplest case",
			target:   7,
			options:  []int{2, 4},
			expected: false,
		},
		{
			desc:     "Another big case",
			target:   10,
			options:  []int{2, 1, 1, 5, 4},
			expected: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := CanSumBruteForce(tC.target, tC.options)
			if res != tC.expected {
				t.Errorf("Expected: %t, but got: %t", tC.expected, res)
			}
		})
	}
}
