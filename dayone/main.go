package main

// it is a O(nˆ2)
func Func_for_test(arr []int, kvalue int) bool {
	for _, value := range arr {
		if value > kvalue {
			continue
		}
		for _, value2 := range arr {
			sum := value + value2
			if sum == kvalue {
				return true
			}
		}
	}
	return false
}

// unsorted array, which contains numbers
// I have a certain

func main() {

}
