package mathskills

import "testing"

func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"Basic average", []int{1, 2, 3, 4, 5}, 3},
		{"Single number", []int{42}, 42},
		{"Two numbers", []int{10, 20}, 15},
		{"Empty slice", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Average(tt.numbers)
			if result != tt.expected {
				t.Errorf("Average(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}