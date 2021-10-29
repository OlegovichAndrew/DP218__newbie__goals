package main

import "testing"

func TestFibo(t *testing.T) {
	testTable := []struct {
		low      int
		high     int
		expected string
	}{
		{
			low:      1,
			high:     5000,
			expected: "1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181",
		},
		{
			low:      0,
			high:     300,
			expected: "0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233",
		},
		{
			low:      200,
			high:     3000,
			expected: "233, 377, 610, 987, 1597, 2584",
		},
	}
	for _, testCase := range testTable {
		result := fibo(testCase.low, testCase.high)
		t.Logf("Runnig with low: %d high: %d, got %s", testCase.low, testCase.high, result)
		if result != testCase.expected {
			t.Errorf("Expected '%s' , got '%s'", testCase.expected, result)
		}
	}
}
