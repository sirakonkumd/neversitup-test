package test

import (
	"neversitup-test/services/permutationService"
	"sort"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "a",
			input:    "a",
			expected: []string{"a"},
		},
		{
			name:     "ab",
			input:    "ab",
			expected: []string{"ab", "ba"},
		},
		{
			name:     "abc",
			input:    "abc",
			expected: []string{"abc", "acb", "bac", "bca", "cba", "cab"},
		},
		{
			name:     "abcc",
			input:    "abcc",
			expected: []string{"bcac", "bcca", "cabc", "cacb", "ccba", "acbc", "accb", "cbac", "cbca", "ccab", "abcc", "bacc"},
		},
	}
	for _, te := range tests {
		t.Run(te.name, func(t *testing.T) {
			res := permutationService.Permutations(te.input)
			sort.Strings(res)
			sort.Strings(te.expected)
			if !equalArr(res, te.expected) {
				t.Errorf("findPermutations(%v) = %v; expected %v", te.input, res, te.expected)
			}
		})
	}
}

func equalArr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
