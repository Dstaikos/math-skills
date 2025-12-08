package mathskills

import "math"

func Sdeviation(numbers []int) int {
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

	variance := s / float64(len(numbers))

	sd := math.Sqrt(variance)

	return int(math.Round(sd))
}
