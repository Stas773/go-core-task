package main

import (
	"reflect"
	"testing"
)

func TestSliceIntersection(t *testing.T) {
	tests := []struct {
		name           string
		slice1, slice2 []int
		expectedBool   bool
		expectedSlice  []int
	}{
		{
			name:          "No intersections",
			slice1:        []int{1, 2, 3},
			slice2:        []int{4, 5, 6},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "Single intersection",
			slice1:        []int{1, 2, 3},
			slice2:        []int{3, 4, 5},
			expectedBool:  true,
			expectedSlice: []int{3},
		},
		{
			name:          "Multiple intersections",
			slice1:        []int{1, 2, 3, 4},
			slice2:        []int{3, 4, 5, 6},
			expectedBool:  true,
			expectedSlice: []int{3, 4},
		},
		{
			name:          "Empty slices",
			slice1:        []int{},
			slice2:        []int{},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "One empty slice",
			slice1:        []int{1, 2, 3},
			slice2:        []int{},
			expectedBool:  false,
			expectedSlice: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			exists, resultSlice := sliceIntersection(test.slice1, test.slice2)
			if exists != test.expectedBool {
				t.Errorf("expected bool %v, got %v", test.expectedBool, exists)
			}
			if !reflect.DeepEqual(resultSlice, test.expectedSlice) {
				t.Errorf("expected value %v, got value %v", test.expectedSlice, resultSlice)
			}
		})
	}
}
