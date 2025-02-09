package main

import (
	"reflect"
	"testing"
)

func TestPrintType(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{42, "int"},
		{052, "int"},
		{0x2A, "int"},
		{3.14, "float64"},
		{"Golang", "string"},
		{true, "bool"},
		{complex64(1 + 2i), "complex64"},
	}

	for _, test := range tests {
		result := printType(test.input)
		if result != test.expected {
			t.Errorf("printType(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected string
	}{
		{[]interface{}{42, 052, 0x2A, 3.14, "Golang", true, 1 + 2i}, "4242423.14Golangtrue(1+2i)"},
		{[]interface{}{1, 2, 3, 4, 5}, "12345"},
		{[]interface{}{"Hellow", " ", "World"}, "Hellow World"},
	}

	for _, test := range tests {
		result := toString(test.input...)
		if result != test.expected {
			t.Errorf("toString(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestStringToRune(t *testing.T) {
	tests := []struct {
		input    string
		expected []rune
	}{
		{"4242423.14Golangtrue(1+2i)", []rune{52, 50, 52, 50, 52, 50, 51, 46, 49, 52, 71, 111, 108, 97, 110, 103, 116, 114, 117, 101, 40, 49, 43, 50, 105, 41}},
		{"Hello", []rune{'H', 'e', 'l', 'l', 'o'}},
		{"12345", []rune{'1', '2', '3', '4', '5'}},
	}
	for _, test := range tests {
		result := stringToRune(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("stringToRune(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestHashRunes(t *testing.T) {
	tests := []struct {
		input    []rune
		salt     string
		expected string
	}{
		{
			[]rune{52, 50, 52, 50, 52, 50, 51, 46, 49, 52, 71, 111, 108, 97, 110, 103, 116, 114, 117, 101, 40, 49, 43, 50, 105, 41},
			"go-2024",
			"53f2f60ac6c41389d3ed3d84d88d8c2860bf8981c677be18243a6f35a6b6a1b3",
		},
		{
			[]rune{'H', 'e', 'l', 'l', 'o', 'w'},
			"salt",
			"f7e52bd801df3a7a8d59798191c63a4e7669cd45a633c6a9c8dc6e6d3b55ff88",
		},
	}
	for _, test := range tests {
		result := hashRunes(test.input, test.salt)
		if result != test.expected {
			t.Errorf("hashRunes(%v, salt: %v) = %v, expected %v", test.input, test.salt, result, test.expected)
		}
	}
}
