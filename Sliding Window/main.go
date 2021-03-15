package main

func CalcAverages(SlideSize int, InputArray []int) []float64 {
	resultArray := make([]float64, 0)
	windowStart := 0
	windowSum := 0.0
	for windowEnd, value := range InputArray {
		windowSum += float64(value)
		if windowEnd >= SlideSize-1 {
			resultArray = append(resultArray, windowSum/float64(SlideSize)) // calculate the average
			windowSum -= float64(InputArray[windowStart])
			windowStart++
		}
	}

	return resultArray
}

func main() {

}
