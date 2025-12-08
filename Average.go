package mathskills

import "math"

func Average(numbers []int) int {
	sum := 0

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	average := float64(sum) / float64(len(numbers))

	return int(math.Round(average))
}
