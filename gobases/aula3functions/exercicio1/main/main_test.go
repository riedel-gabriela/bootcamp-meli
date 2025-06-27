package main

import (
	"testing"
)

func TestImpostoSobreSalario(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{49000.0, 0},
		{51000.0, 16830.0},
		{151000.0, 25770.0},
	}

	for _, tc := range testCases {
		actual := impostoSobreSalario(tc.input)
		if actual != tc.expected {
			t.Errorf("impostoSobreSalatio(%f) = %f; want %f", tc.input, actual, tc.expected)
		}
	}

}
