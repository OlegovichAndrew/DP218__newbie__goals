package main

import "testing"

func TestReadAndCount(t *testing.T) {
	testTable := []struct {
		file      string
		forSearch string
		expected  int
	}{
		{
			file:      "file.txt",
			forSearch: "pussy",
			expected:  13,
		},
		{
			file:      "file.txt",
			forSearch: "believe",
			expected:  3,
		},
		{
			file:      "file.txt",
			forSearch: "not",
			expected:  2,
		},
	}

	for _, testCase := range testTable {
		result, _ := readAndCount(testCase.file, testCase.forSearch)
		t.Logf("Calling func with: %s, got result: %v ", testCase.forSearch, result)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}

func TestRewriteString(t *testing.T) {
	testTable := []struct {
		file           string
		textFind       string
		textWrite      string
		expectedPussy  int
		expectedPussy2 int
	}{
		{
			file:           "file.txt",
			textFind:       "pussy",
			textWrite:      "pussy2",
			expectedPussy2: 13,
		},
	}

	for _, testCase := range testTable {
		t.Logf("Calling func with: %s, and: %v ", testCase.textFind, testCase.textWrite)
		rewriteString(testCase.file, testCase.textFind, testCase.textWrite)
		pussy2, _ := readAndCount(testCase.file, testCase.textWrite)

		if pussy2 != testCase.expectedPussy2 {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expectedPussy2, pussy2)
		}
	}
}
