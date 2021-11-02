package main

import "testing"

func TestCreateBoard(t *testing.T) {
	var expected = [][]string{
		{},
		{whiteSquare},
		{whiteSquare + blackSquare, blackSquare + whiteSquare},
	}
	var results = [][]string{
		createBoard(0, 0),
		createBoard(1, 1),
		createBoard(2, 2),
	}

	for i := 0; i < len(expected); i++ {
		outExpect := expected[i]
		outResult := results[i]

		if len(outExpect) != len(outResult) {
			t.Errorf("Error: expected: %d , received %d", len(outExpect), len(outResult))
		}
		for j, v := range outResult {
			if outResult[j] != v {
				t.Errorf("Error %d expected %v, got %v", i, outExpect[i], v)
			}
		}
	}
}
