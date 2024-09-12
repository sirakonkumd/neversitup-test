package test

import (
	"neversitup-test/services/smileService"
	"testing"
)

func TestSmile(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "return 3",
			input:    []string{":)", ";(", ";}", ":D", ":-()", "(:-)", ";~)"},
			expected: 3,
		},
		{
			name:     "return 2",
			input:    []string{":)", ";(", ";}", ":-D"},
			expected: 2,
		},
		{
			name:     "return 2",
			input:    []string{":)", ";(", ";}", ":{", ":-()", "(:-)"},
			expected: 1,
		},
		{
			name:     "return 3",
			input:    []string{";D", ":-(", ":-)", ";~)"},
			expected: 3,
		},
		{
			name:     "return 3",
			input:    []string{":)", ";(", ";}", ":D", ":-()", "(:-)", ";~)"},
			expected: 3,
		},
	}

	for _, te := range tests {
		t.Run(te.name, func(t *testing.T) {
			res := smileService.FindSmile(te.input)
			if res != te.expected {
				t.Errorf("findSmile(%v) = %v; expected %v", te.input, res, te.expected)
			}
		})
	}
}
