package mathskills

import "math"

func Median(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	
	nums := append([]int{}, numbers...) // making a copy of the slice

	// bubblesort
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	if len(nums)%2 != 0 {
		return nums[len(nums)/2]
	}

	a := nums[len(nums)/2-1]
	b := nums[len(nums)/2]

	median := float64(a+b) / 2.0

	return int(math.Round(median))
}
