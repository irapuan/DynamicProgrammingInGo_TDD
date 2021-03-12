package main

import "testing"

func Test(t *testing.T) {
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
