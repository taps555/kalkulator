package main

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		a, b      float64
		operator  string
		expected  float64
		shouldErr bool
	}{
		{10, 5, "+", 15, false},
		{10, 5, "-", 5, false},
		{10, 5, "*", 50, false},
		{10, 5, "/", 2, false},
		{10, 0, "/", 0, true},           // Error: division by zero
		{10, 5, "%", 0, true},           // Error: invalid operator
		{0, 0, "/", 0, true},            // Error: division by zero
		{-10, 5, "+", -5, false},        // Negative numbers
		{-10, -5, "*", 50, false},       // Both negative numbers
		{1.5, 2.5, "+", 4, false},       // Floating-point numbers
	}

	for _, test := range tests {
		result, err := Calculate(test.a, test.b, test.operator)
		if test.shouldErr {
			if err == nil {
				t.Errorf("Expected error for input (%v, %v, %v), but got nil", test.a, test.b, test.operator)
			}
		} else {
			if err != nil {
				t.Errorf("Did not expect error for input (%v, %v, %v), but got: %v", test.a, test.b, test.operator, err)
			}
			if result != test.expected {
				t.Errorf("For input (%v, %v, %v), expected %v but got %v", test.a, test.b, test.operator, test.expected, result)
			}
		}
	}
}
