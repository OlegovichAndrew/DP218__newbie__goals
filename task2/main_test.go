package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCompareLetters(t *testing.T) {

	testTable := []struct {
		letter1  Letter
		letter2  Letter
		expected string
	}{
		{
			letter1:  Letter{15, 14},
			letter2:  Letter{14, 15},
			expected: "Both letters are equal!",
		},
		{
			letter1:  Letter{22, 20},
			letter2:  Letter{21, 25},
			expected: "You CAN put the first letter into the second",
		},
		{
			letter1:  Letter{34, 22},
			letter2:  Letter{19, 30},
			expected: "You CAN put the second letter into the first",
		},
		{
			letter1:  Letter{14, 13},
			letter2:  Letter{42, 11},
			expected: "You CAN'T put letters into each other",
		},
	}

	for _, testCase := range testTable {
		result := compareLetters(testCase.letter1, testCase.letter2)
		t.Logf("Calling %v and %v, got %v", testCase.letter1, testCase.letter2, result)

		if result != testCase.expected {
			t.Errorf("Incorrect result, expected %s, got %s", testCase.expected, result)
		}
	}
}

func TestAnalyzeAnswer(t *testing.T) {
	testTable := []struct {
		answer   string
		expected bool
	}{
		{
			answer:   "YeS",
			expected: true,
		},
		{
			answer:   "Y",
			expected: true,
		},
		{
			answer:   "something",
			expected: false,
		},
		{
			answer:   "",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		result := analyzeAnswer(testCase.answer)
		t.Logf("Calling %v , got %v", testCase.answer, result)
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}

func TestAskUserValues(t *testing.T) {
	testTable := []struct {
		input          string
		expectedSideA1 float64
		expectedSideB1 float64
		expectedSideA2 float64
		expectedSideB2 float64
	}{
		{
			input:          `this\n word\n and\n this`,
			expectedSideA1: 0,
			expectedSideB1: 0,
			expectedSideA2: 0,
			expectedSideB2: 0,
		},
		{
			input:          "22\n34\n53\n48",
			expectedSideA1: 22,
			expectedSideB1: 34,
			expectedSideA2: 53,
			expectedSideB2: 48,
		},
	}

	for _, testCase := range testTable {
		content := []byte(testCase.input)
		tmpFile, err := ioutil.TempFile("", "example")
		if err != nil {
			t.Error(err)
		}

		if _, err := tmpFile.Write(content); err != nil {
			t.Error(err)
		}
		if _, err := tmpFile.Seek(0, 0); err != nil {
			t.Error(err)
		}

		oldStdin := os.Stdin

		os.Stdin = tmpFile

		if askUserValues(); letter1.SideA != testCase.expectedSideA1 {
			t.Errorf("AskUserValues failed: value - %v, expected %v", letter1.SideA, testCase.expectedSideA1)
		}
		if askUserValues(); letter1.SideB != testCase.expectedSideB1 {
			t.Errorf("AskUserValues failed: value - %v, expected %v", letter1.SideB, testCase.expectedSideB1)
		}
		if askUserValues(); letter2.SideA != testCase.expectedSideA2 {
			t.Errorf("AskUserValues failed: value - %v, expected %v", letter2.SideA, testCase.expectedSideA2)
		}
		if askUserValues(); letter2.SideB != testCase.expectedSideB2 {
			t.Errorf("AskUserValues failed: value - %v, expected %v", letter2.SideB, testCase.expectedSideB2)
		}

		if err := tmpFile.Close(); err != nil {
			log.Fatal(err)
		}

		os.Remove(tmpFile.Name())
		func() { os.Stdin = oldStdin }()
	}
}
