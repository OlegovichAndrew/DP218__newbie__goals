package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
	var scanner = bufio.NewScanner(os.Stdin)

	for {
	START1:
		fmt.Println("Enter 'a' side length of the 1-st letter")
		scanner.Scan()
		letter1.SideA, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			goto START1
		}

	START2:
		fmt.Println("Enter 'b' side length of the 1-st letter")
		scanner.Scan()
		letter1.SideB, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			goto START2
		}

	START3:
		fmt.Println("Enter 'c' side length of the 2-nd letter")
		scanner.Scan()
		letter2.SideA, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			goto START3
		}
	START4:
		fmt.Println("Enter 'd' side length of the 2-nd letter")
		scanner.Scan()
		letter2.SideB, err = strconv.ParseFloat(scanner.Text(), 64)
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
	fmt.Scanf("%s", &answer)
	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		main()
	}
	return
}

func main() {
	askUserValues()
	compareLetters()
	askUserContinue()
}
