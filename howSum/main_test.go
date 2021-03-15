package howsum

import (
	"reflect"
	"testing"
)

func TestHowSumBruteForce(t *testing.T) {
	testCases := []struct {
		desc      string
		targetsum int
		inputs    []int
		expected  []int
	}{
		{
			desc:      "Simplest case",
			targetsum: 1,
			inputs:    []int{},
			expected:  nil,
		},
		{
			desc:      "Simplest case 2",
			targetsum: 1,
			inputs:    []int{1},
			expected:  []int{1},
		},
		{
			desc:      "Simplest case 2",
			targetsum: 2,
			inputs:    []int{1},
			expected:  nil,
		},
		{
			desc:      "Simplest case 3",
			targetsum: 2,
			inputs:    []int{1, 1},
			expected:  []int{1, 1},
		},
		{
			desc:      "Simplest case 4",
			targetsum: 5,
			inputs:    []int{1, 1, 1, 2},
			expected:  []int{1, 1, 1, 2},
		},
		{
			desc:      "Simplest case 5",
			targetsum: 8,
			inputs:    []int{1, 1, 1, 2},
			expected:  nil,
		},
		{
			desc:      "Simplest case 5",
			targetsum: 6,
			inputs:    []int{1, 4, 9, 1},
			expected:  []int{1, 4, 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := HowSum(tC.targetsum, tC.inputs)
			if !reflect.DeepEqual(res, tC.expected) {
				t.Errorf("Expected: %v , but got %v", tC.expected, res)
			}
		})
	}
}
