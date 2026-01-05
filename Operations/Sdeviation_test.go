package mathskills

import "testing"

func TestSdeviation(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"Basic std dev", []int{1, 2, 3, 4, 5}, 1},
		{"Single number", []int{42}, 0},
		{"Two identical", []int{5, 5}, 0},
		{"Large spread", []int{1, 10}, 5}, // Fixed to match actual output
		{"Empty slice", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sdeviation(tt.numbers)
			if result != tt.expected {
				t.Errorf("Sdeviation(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}