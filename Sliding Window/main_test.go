package main

import (
	"reflect"
	"testing"
)

func TestSlidingWindow(t *testing.T) {
	testCases := []struct {
		desc       string
		inputarray []int
		slidesize  int
		expected   []float64
	}{
		{
			desc:       "First test",
			inputarray: []int{1, 3, 2, 6, -1, 4, 1, 8, 2},
			slidesize:  5,
			expected:   []float64{2.2, 2.8, 2.4, 3.6, 2.8},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := CalcAverages(tC.slidesize, tC.inputarray)
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("Expected: %v, but got %v", tC.expected, result)
			}
		})
	}
}
