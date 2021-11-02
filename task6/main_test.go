package main

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	filePath := "moscow.txt"
	expected := "moscow"

	result := readFile(filePath)

	if result != expected {
		t.Errorf("Incorrect result. Expect %v, got %v", expected, result)
	}
}

func TestPiterAlgorythm(t *testing.T) {
	testTable := []struct {
		ticket   string
		expected bool
	}{
		{
			ticket:   "044530",
			expected: true,
		},
		{
			ticket:   "356536",
			expected: false,
		},
		{
			ticket:   "000001",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		result := piterAlgorythm(testCase.ticket)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}

func TessMoscowAlgorythm(t *testing.T) {
	testTable := []struct {
		ticket   string
		expected bool
	}{
		{
			ticket:   "044530",
			expected: false,
		},
		{
			ticket:   "356536",
			expected: true,
		},
		{
			ticket:   "000001",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		result := moscowAlgorythm(testCase.ticket)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}

func Test_Main(t *testing.T) {
	osArgsPrevious := os.Args
	defer func() { os.Args = osArgsPrevious }()
	os.Args = []string{"main.go", "piter.txt"}
	main()
	os.Args = []string{"main.go", "moscow.txt"}
	main()
}
