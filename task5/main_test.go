package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestInputOperations(t *testing.T) {
	testTable := []struct {
		input       string
		expected    []string
		expectedErr error
	}{
		{
			input:       "2223fdsf232",
			expected:    []string{},
			expectedErr: errors.New("your input is not correct"),
		},
		{
			input:       "2223232543543543321",
			expected:    []string{},
			expectedErr: errors.New("cannot handle your number"),
		},
		{
			input:       "32421",
			expected:    []string{"000", "000", "000", "32", "421"},
			expectedErr: nil,
		},
		{
			input:       "43253452321",
			expected:    []string{"000", "43", "253", "452", "321"},
			expectedErr: nil,
		},
		{
			input:       "0",
			expected:    []string{"000", "000", "000", "000", "0"},
			expectedErr: nil,
		},
		{
			input:       "00",
			expected:    []string{"000", "000", "000", "000", "00"},
			expectedErr: nil,
		},
	}

	for _, testCase := range testTable {
		result, err := inputOperations(testCase.input)
		t.Logf("Calling func with: %s, got result: %v ", testCase.input, result)

		if !reflect.DeepEqual(err, testCase.expectedErr) {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expectedErr, err)
		}
		if !reflect.DeepEqual(err, testCase.expectedErr) {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}

func TestAddToThree(t *testing.T) {
	testTable := []struct {
		input    string
		expected string
	}{
		{
			input:    "2",
			expected: "002",
		},
		{
			input:    "31",
			expected: "031",
		},
		{
			input:    "564",
			expected: "564",
		},
	}

	for _, testCase := range testTable {
		result := addToThree(testCase.input)
		t.Logf("Calling func with: %s, got result: %v ", testCase.input, result)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}

func TestGivePieceNames(t *testing.T) {
	testTable := []struct {
		input    string
		expected string
	}{
		{
			input:    "312",
			expected: "триста двенадцать ",
		},
		{
			input:    "534",
			expected: "пятьсот тридцать четыре ",
		},
		{
			input:    "05",
			expected: "пять ",
		},
		{
			input:    "013",
			expected: "тринадцать ",
		},
	}

	for _, testCase := range testTable {
		result := givePieceNames(testCase.input)
		t.Logf("Calling func with: %s, got result: %v ", testCase.input, result)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}

func TestGiveThousandPiecesNames(t *testing.T) {

	testTable := []struct {
		input    string
		expected string
	}{
		{
			input:    "312",
			expected: "триста двенадцать ",
		},
		{
			input:    "534",
			expected: "пятьсот тридцать четыре ",
		},
		{
			input:    "05",
			expected: "пять ",
		},
		{
			input:    "013",
			expected: "тринадцать ",
		},
		{
			input:    "022",
			expected: "двадцать две ",
		},
		{
			input:    "31",
			expected: "тридцать одна ",
		},
	}

	for _, testCase := range testTable {
		result := giveThousandPiecesNames(testCase.input)
		t.Logf("Calling func with: %s, got result: %v ", testCase.input, result)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}

func TestAddEnding(t *testing.T) {
	testTable := []struct {
		input    string
		expected string
	}{
		{
			input:    "312",
			expected: "ов ",
		},
		{
			input:    "342",
			expected: "а ",
		},
		{
			input:    "551",
			expected: " ",
		},
	}

	for _, testCase := range testTable {
		result := addEnding(testCase.input)
		t.Logf("Calling func with: %s, got result: %v ", testCase.input, result)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)

		}
	}
}

func TestAddEndingThousand(t *testing.T) {
	testTable := []struct {
		input    string
		expected string
	}{
		{
			input:    "361",
			expected: "а ",
		},
		{
			input:    "342",
			expected: "и ",
		},
		{
			input:    "558",
			expected: " ",
		},
	}

	for _, testCase := range testTable {
		result := addEndingThousand(testCase.input)
		t.Logf("Calling func with: %s, got result: %v ", testCase.input, result)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)

		}
	}
}

func TestCompletePieces(t *testing.T) {
	testTable := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"000", "43", "253", "452", "321"},
			expected: "сорок три миллиарда двести пятьдесят три миллиона четыреста пятьдесят две тысячи триста двадцать один ",
		},
		{
			input:    []string{"000", "1", "000", "000", "001"},
			expected: "один миллиард один ",
		},
		{
			input:    []string{"000", "000", "000", "000", "000"},
			expected: "ноль",
		},
		{
			input:    []string{"012", "334", "222", "731", "500"},
			expected: "двенадцать триллионов триста тридцать четыре миллиарда двести двадцать два миллиона семьсот тридцать одна тысяча пятьсот ",
		},
	}
	for _, testCase := range testTable {
		result := completePieces(testCase.input)
		t.Logf("Calling func with: %s, got result: %v ", testCase.input, result)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}

}
