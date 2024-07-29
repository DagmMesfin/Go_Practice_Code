package main

import (
	"reflect"
	"testing"
)

type TestFreq struct {
	input    string
	expected map[string]int
}

func TestWordFrequency(t *testing.T) {
	tests := []TestFreq{
		{
			input:    "Hello, world! Hello world?",
			expected: map[string]int{"hello": 2, "world": 2},
		},
		{
			input:    "It's a wonderful world, indeed. Wonderful!",
			expected: map[string]int{"its": 1, "a": 1, "wonderful": 2, "world": 1, "indeed": 1},
		},
		{
			input:    "Go is awesome. Go Go Go!",
			expected: map[string]int{"go": 4, "is": 1, "awesome": 1},
		},
		{
			input:    "Empty,,,      spaces!!",
			expected: map[string]int{"empty": 1, "spaces": 1},
		},
		{
			input:    "",
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		result := wordCount(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %q, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

type TestPal struct {
	input    string
	expected bool
}

func TestPalindrome(t *testing.T) {
	tests := []TestPal{
		{
			input:    "racecar",
			expected: true,
		},
		{
			input:    "level",
			expected: true,
		},
		{
			input:    "apple",
			expected: false,
		},
		{
			input:    "banana",
			expected: false,
		},
		{
			input:    "kayak",
			expected: true,
		},
	}

	for _, test := range tests {
		result := palindromeChecker(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
