package main

import (
	"testing"
)

type TestFreq struct {
	input          map[string]float32
	expected_avg   float32
	expected_grade string
}

func TestGrade(t *testing.T) {
	tests := []TestFreq{
		{
			input: map[string]float32{
				"Math":    85,
				"English": 90,
				"Science": 78,
				"History": 65,
				"Art":     50,
			},
			expected_avg:   float32(73.6),
			expected_grade: "B",
		},
	}

	for num, test := range tests {
		result_avg := avgCalculator(test.input)
		result_grade := markGrade(result_avg)
		if result_avg != test.expected_avg && result_grade != test.expected_grade {
			t.Errorf("For input %d, expected %v, but got %v", num, test.expected_avg, result_avg)
			t.Errorf("For input %d, expected %v, but got %v", num, test.expected_grade, result_grade)
		}
	}
}
