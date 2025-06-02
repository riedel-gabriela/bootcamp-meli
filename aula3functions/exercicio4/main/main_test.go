package main

import "testing"

func TestMinValue(t *testing.T) {
	expectedResult := 0

	result := minValue(0, 1, 2, 3, 4)

	if result != float64(expectedResult) {
		t.Errorf("Expected %d but got %fd", expectedResult, result)
	}
}

func TestMaxValue(t *testing.T) {
	expectedResult := 4

	result := maxValue(0, 1, 2, 3, 4)

	if result != float64(expectedResult) {
		t.Errorf("Expected %d but got %fd", expectedResult, result)
	}
}

func TestAverageValue(t *testing.T) {
	expectedResult := 2

	result := averageValue(0, 1, 2, 3, 4)

	if result != float64(expectedResult) {
		t.Errorf("Expected %d but got %fd", expectedResult, result)
	}
}

func TestOperation(t *testing.T) {
	tests := []struct {
		name         string
		opType       string
		expectedFunc func(numbers ...float64) float64
	}{
		{
			name:         "Minimum operation",
			opType:       minimum,
			expectedFunc: minValue,
		},
		{
			name:         "Maximum operation",
			opType:       maximum,
			expectedFunc: maxValue,
		},
		{
			name:         "Average operation",
			opType:       average,
			expectedFunc: averageValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn, err := operation(tt.opType)
			if err != nil {
				t.Fatalf("operation returned an error for %s: %v", tt.opType, err)
			}
			if fn == nil {
				t.Errorf("operation returned nil function for %s", tt.opType)
			}
		})
	}
}
