package main

import (
	"reflect"
	"testing"
)

func TestSliceDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "No common elements",
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{"date", "fig", "grape"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "Some common elements",
			slice1:   []string{"apple", "banana", "cherry", "date"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry"},
		},
		{
			name:     "All elements common",
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{"apple", "banana", "cherry"},
			expected: []string{},
		},
		{
			name:     "Empty first slice",
			slice1:   []string{},
			slice2:   []string{"apple", "banana"},
			expected: []string{},
		},
		{
			name:     "Empty second slice",
			slice1:   []string{"apple", "banana"},
			slice2:   []string{},
			expected: []string{"apple", "banana"},
		},
		{
			name:     "Both slices empty",
			slice1:   []string{},
			slice2:   []string{},
			expected: []string{},
		},
		{
			name:     "Duplicates in first slice",
			slice1:   []string{"apple", "banana", "apple", "cherry"},
			slice2:   []string{"banana", "date"},
			expected: []string{"apple", "apple", "cherry"},
		},
		{
			name:     "Duplicates in second slice",
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{"banana", "banana", "date"},
			expected: []string{"apple", "cherry"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultSlice := sliceDifference(test.slice1, test.slice2)

			if !reflect.DeepEqual(resultSlice, test.expected) {
				t.Errorf("expected value %v, got value %v", test.expected, resultSlice)
			}
		})
	}
}
