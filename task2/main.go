package main

import (
	"fmt"
	"math"
	"strings"
)

type Letter struct {
	SideA, SideB float64
}

var letter1 Letter
var letter2 Letter
var answer string
var err error

func askUserValues() {

	for {
START1:
		fmt.Println("Enter 'a' side of the 1-st letter")
		_, err = fmt.Scan(&letter1.SideA)
		if err != nil {
			goto START1
		}
START2:
		fmt.Println("Enter 'b' side of  the 1-st letter")
		_, err = fmt.Scan(&letter1.SideB)
		if err != nil {
			goto START2
		}
START3:
		fmt.Println("Enter 'c' side of the 2-nd letter")
		_, err = fmt.Scan(&letter2.SideA)
		if err != nil {
			goto START3
		}
START4:
		fmt.Println("Enter 'd' side of the 2-nd letter")
		_, err = fmt.Scan(&letter2.SideB)
		if err != nil {
			goto START4
		}
		return
	}
}

func compareLetters() {
	if math.Min(letter1.SideA, letter1.SideB) == math.Min(letter2.SideA, letter2.SideB) && math.Max(letter1.SideA, letter1.SideB) == math.Max(letter2.SideA, letter2.SideB) {
		fmt.Println("Both letters are equal!")
	} else if math.Min(letter1.SideA, letter1.SideB) >= math.Min(letter2.SideA, letter2.SideB) && math.Max(letter1.SideA, letter1.SideB) >= math.Max(letter2.SideA, letter2.SideB) {
		fmt.Println("You CAN put the second letter into the first")
	} else if math.Min(letter1.SideA, letter1.SideB) <= math.Min(letter2.SideA, letter2.SideB) && math.Max(letter1.SideA, letter1.SideB) <= math.Max(letter2.SideA, letter2.SideB) {
		fmt.Println("You CAN put the first letter into the second")
	} else {
		fmt.Println(`You CAN'T put letters into each other`)
	}
}

func askUserContinue() {
	fmt.Println(`Would you like to continue the comparison? Enter "y" or "yes" to continue`)
	fmt.Scanf("%v", &answer)
	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		askUserValues()
	}
	return
}

func main() {
	askUserValues()
	compareLetters()
	askUserContinue()
}