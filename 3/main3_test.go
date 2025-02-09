package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		key      string
		value    int
		expected int
	}{
		{"first", 1, 1},
		{"new key", 125, 125},
		{"", 0, 0},
	}
	for _, test := range tests {
		m := NewStringIntMap()
		m.Add(test.key, test.value)
		if m.db[test.key] != test.expected {
			t.Errorf("For key %s: expected value %d, got value %d", test.key, test.expected, test.value)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		key       string
		initialDB map[string]int
		expected  bool
	}{
		{"first", map[string]int{"first": 1, "second": 2}, false},
		{"non exist key", map[string]int{"first": 1, "second": 2}, false},
		{"first", map[string]int{"second": 2}, false},
	}
	for _, test := range tests {
		m := NewStringIntMap()
		for key, v := range test.initialDB {
			m.db[key] = v
		}
		m.Remove(test.key)
		if _, ok := m.db[test.key]; ok != test.expected {
			t.Errorf("For key %s: expected value %v, got value %v", test.key, test.expected, ok)
		}
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		initialDB map[string]int
		expected  map[string]int
	}{
		{map[string]int{"first": 1, "second": 2}, map[string]int{"first": 1, "second": 2}},
		{map[string]int{"one element": 4}, map[string]int{"one element": 4}},
		{map[string]int{}, map[string]int{}},
	}
	for _, test := range tests {
		m := NewStringIntMap()
		for key, v := range test.initialDB {
			m.db[key] = v
		}
		m.Copy()
		if !reflect.DeepEqual(m.db, test.expected) {
			t.Errorf("expected value %v, got value %v", test.expected, m.db)
		}
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		key       string
		initialDB map[string]int
		expected  bool
	}{
		{"first", map[string]int{"first": 1, "second": 2}, true},
		{"non exist key", map[string]int{"first": 1, "second": 2}, false},
		{"first", map[string]int{}, false},
	}
	for _, test := range tests {
		m := NewStringIntMap()
		for key, v := range test.initialDB {
			m.db[key] = v
		}
		if ok := m.Exists(test.key); ok != test.expected {
			t.Errorf("For key %s: expected value %v, got value %v", test.key, test.expected, ok)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		key       string
		initialDB map[string]int
		value     int
		expected  bool
	}{
		{"first", map[string]int{"first": 1, "second": 2}, 1, true},
		{"non exist key", map[string]int{"first": 1, "second": 2}, 0, false},
		{"first", map[string]int{}, 0, false},
	}
	for _, test := range tests {
		m := NewStringIntMap()
		for key, v := range test.initialDB {
			m.db[key] = v
		}
		if v, ok := m.Get(test.key); ok != test.expected || v != test.value {
			t.Errorf("For key %s: expected %v and value %v, got %v and value %v", test.key, test.expected, test.value, ok, v)
		}
	}
}
