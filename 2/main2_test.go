package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{2, 4, 6, 8, 10}},
		{[]int{1, 3, 5, 7}, nil},
	}
	for _, test := range tests {
		result := sliceExample(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("sliceExample(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		input    []int
		element  int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 156643, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 156643}},
	}
	for _, test := range tests {
		result := addElements(test.input, test.element)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("addElements(%v, %v) = %v, expected %v", test.input, test.element, result, test.expected)
		}
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{}, []int{}},
	}
	for _, test := range tests {
		result := copySlice(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("addElements(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		input    []int
		index    int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, []int{1, 2, 3, 4, 5, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, test := range tests {
		result := removeElements(test.input, test.index)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("removeElements(%v, %v) = %v, expected %v", test.input, test.index, result, test.expected)
		}
	}
}
