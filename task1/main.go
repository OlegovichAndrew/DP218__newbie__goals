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
	flag.IntVar(&height, "h",0, "The height (int) of the chessboard" )
	flag.IntVar(&width, "w", 0, "The width (int) of the chessboard")
	flag.Parse()
	isInputOk := (height > 0 && width > 0)
	return isInputOk
}

func createBoard()  {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i % 2 == 0 && j % 2 == 0 || i % 2 != 0 && j % 2 != 0 {
				fmt.Printf("%s", whiteSquare)
			} else  {
				fmt.Printf("%s", blackSquare)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	if inputParams() {
		createBoard()§
	} else {
		fmt.Println(instruction)
	}
}




