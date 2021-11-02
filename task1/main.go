package main

import (
	"flag"
	"fmt"
)

const (
	blackSquare = "•"
	whiteSquare = " "
	instruction = "----------------------------------------------------\n" +
		"You have to input both options 'height' and 'width' \n" +
		"You can do it with flags -h -w \n" +
		"----------------------------------------------------"
)

var (
	width, height int
)

func inputParams() bool {
	flag.IntVar(&height, "h", 0, "The height (int) of the chessboard")
	flag.IntVar(&width, "w", 0, "The width (int) of the chessboard")
	flag.Parse()
	return height > 0 && width > 0
}

func createBoard(width, height int) []string {
	result := make([]string, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i%2 == 0 && j%2 == 0 || i%2 != 0 && j%2 != 0 {
				result[i] = result[i] + whiteSquare
			} else {
				result[i] = result[i] + blackSquare
			}
		}
	}
	return result
}

func main() {
	if inputParams() {
		board := createBoard(width, height)
		for _, v := range board {
			fmt.Println(v)
		}
	} else {
		fmt.Println(instruction)
	}
}
