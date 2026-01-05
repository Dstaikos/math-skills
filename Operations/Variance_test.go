package mathskills

import "testing"

func TestVariance(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"Basic variance", []int{1, 2, 3, 4, 5}, 2},
		{"Single number", []int{42}, 0},
		{"Two identical", []int{5, 5}, 0},
		{"Large spread", []int{1, 10}, 20}, // Fixed expectation
		{"Empty slice", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Variance(tt.numbers)
			if result != tt.expected {
				t.Errorf("Variance(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}