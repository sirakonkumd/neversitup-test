package test

import (
	"neversitup-test/services/oddService"
	"testing"
)

func TestFindOdd(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "return 7",
			input:    []int{7},
			expected: 7,
		},
		{
			name:     "return 0",
			input:    []int{0},
			expected: 0,
		},
		{
			name:     "return 2",
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "return 0",
			input:    []int{0, 1, 0, 1, 0},
			expected: 0,
		},
		{
			name:     "return 4",
			input:    []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			expected: 4,
		},
		{
			name:     "Empty => 0",
			input:    []int{},
			expected: 0,
		},
		// {
		// 	name:     "fail",
		// 	input:    []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
		// 	expected: 2,
		// },
	}

	for _, te := range tests {
		t.Run(te.name, func(t *testing.T) {
			res := oddService.FindOdd(te.input)
			if res != te.expected {
				t.Errorf("findOdd(%v) = %v; expected %v", te.input, res, te.expected)
			}
		})
	}
}
