package mathskills

import "testing"

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"Odd length", []int{1, 3, 2}, 2},
		{"Even length", []int{1, 2, 3, 4}, 3},
		{"Single number", []int{42}, 42},
		{"Empty slice", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Median(tt.numbers)
			if result != tt.expected {
				t.Errorf("Median(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}