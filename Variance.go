package mathskills

import "math"

func Variance(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	mean := float64(sum) / float64(len(numbers))

	var s float64

	for _, v := range numbers {

		diff := float64(v) - mean
		s += diff * diff
	}

	Variance := s / float64(len(numbers))

	return int(math.Round(Variance))
}
