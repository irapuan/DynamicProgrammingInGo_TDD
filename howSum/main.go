package howsum

func HowSum(target int, options []int) []int {
	if target < 0 {
		return nil
	}
	if len(options) == 1 && options[0] == target {
		return options
	}
	if target == 0 {
		return make([]int, 0)
	}

	for _, value := range options {
		innerResult := HowSum(target-value, options)
		if innerResult != nil {
			result := make([]int, 0)
			result = append(result, value)
			result = append(result, innerResult...)
			return result
		}
	}
	return nil
}

// HowSum(2,[1,1]) = [1,1]
// HowSum(1,[1]) = [1]
// HowSum(1,[2]) = nil
// HowSum(3,[1,1,1]) => append( HowSum(1,[1]), HowSum(2,[1,1]) ) => append( HowSum(1,[1]), HowSum ])
// HowSum(4,[1,1,2]) => append( HowSum(1,[1]), HowSum(3,[1,2]) )
