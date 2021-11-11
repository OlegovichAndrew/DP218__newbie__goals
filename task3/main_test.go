package main

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func TestUserInput(t *testing.T) {
	input := []byte("Example, 44, 55, 86")
	expected := string(input)
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(input)
	if err != nil {
		t.Fatal(err)
	}
	w.Close()

	stdin := os.Stdin

	defer func() { os.Stdin = stdin }()
	os.Stdin = r

	if expected != userInput(os.Stdin) {
		t.Fatalf("userInput: %v", userInput(os.Stdin))
	}
}

func TestPrepareInput(t *testing.T) {
	testTable := []struct {
		input       string
		expectedErr error
	}{
		{
			input:       "case1, 33, 44, 55",
			expectedErr: nil,
		},
		{
			input:       "  case2   , 28,   12 , 19  ",
			expectedErr: nil,
		},
		{
			input:       "case3, 18, 30, 40, 60, 20",
			expectedErr: errors.New("wrong number of arguments\n"),
		},
		{
			input:       "case4, 3, side , 8",
			expectedErr: errors.New("wrong side value\n"),
		},
		{
			input:       "case5, 5, 7, 100",
			expectedErr: errors.New(`can't build a triangle with inputted values`),
		},
	}

	for _, testCase := range testTable {
		_, _, _, _, result := prepareInput(testCase.input)
		t.Logf("Calling func with: %s, got result: %v ", testCase.input, result)
		if !reflect.DeepEqual(result, testCase.expectedErr) {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expectedErr, result)
		}
	}
}

func TestFormulaHeron(t *testing.T) {
	testTable := []struct {
		income       string
		expectedName string
		expectedArea float64
	}{
		{
			income:       "case1, 44, 33, 55",
			expectedName: "case1",
			expectedArea: 726,
		},
		{
			income:       "case2, 44, 55, 64",
			expectedName: "case2",
			expectedArea: 1190.52,
		},
		{
			income:       "case3, 44, side, 20",
			expectedName: "",
			expectedArea: 0,
		},
	}

	for _, testCase := range testTable {
		name, area, _ := formulaHeron(prepareInput(testCase.income))
		t.Logf("Calling with: %s, has %s area: %f", testCase.income, name, area)
		if name != testCase.expectedName {
			t.Errorf("Wanted name: %s, got %s", testCase.expectedName, name)
		}
		if area != testCase.expectedArea {
			t.Errorf("Expected area: %f, got %f", testCase.expectedArea, area)
		}
	}
}

func TestAddToMap(t *testing.T) {
	var testTable = []struct {
		err      error
		name     string
		area     float64
		expected map[string]float64
	}{
		{
			err:      errors.New("some error"),
			name:     "some name",
			area:     500,
			expected: map[string]float64{},
		},
		{
			err:      nil,
			name:     "first",
			area:     450.5,
			expected: map[string]float64{"first": 450.5},
		},
		{
			err:      nil,
			name:     "second",
			area:     245.5454,
			expected: map[string]float64{"first": 450.5, "second": 245.5454},
		},
	}

	for _, testCase := range testTable {
		result := addToMap(testCase.name, testCase.area, testCase.err)
		t.Logf("Fucntion calling with args: %v, %v, %v, got the result: %v", testCase.name, testCase.area, testCase.err, result)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Test failed. Expected: %v, but got: %v", testCase.expected, result)
		}
	}
}

func Test_main(t *testing.T) {
	osArgsPrevious := os.Args
	defer func() { os.Args = osArgsPrevious }()
	os.Args = []string{"TestMain, 30, 35,40"}
	main()
}
