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
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := CountArrayItems(tC.inputlist)
			if res != tC.expected {
				t.Errorf("Expected: %d, but got %d", tC.expected, res)
			}
		})
	}
}
