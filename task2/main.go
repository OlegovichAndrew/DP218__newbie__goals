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
var err error

func askUserValues() {
	var scanner = bufio.NewScanner(os.Stdin)

	for {
	START1:
		fmt.Println("Enter 'a' side length of the 1-st letter")
		if scanner.Scan() {
			letter1.SideA, err = strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				goto START1
			}
		}
	START2:
		fmt.Println("Enter 'b' side length of the 1-st letter")
		if scanner.Scan() {
			letter1.SideB, err = strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				goto START2
			}
		}
	START3:
		fmt.Println("Enter 'c' side length of the 2-nd letter")
		if scanner.Scan() {
			letter2.SideA, err = strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				goto START3
			}
		}

	START4:
		fmt.Println("Enter 'd' side length of the 2-nd letter")
		if scanner.Scan() {
			letter2.SideB, err = strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				goto START4
			}
		}
		return
	}
}

func compareLetters(letter1, letter2 Letter) string {
	switch {
	case math.Min(letter1.SideA, letter1.SideB) == math.Min(letter2.SideA, letter2.SideB) && math.Max(letter1.SideA, letter1.SideB) == math.Max(letter2.SideA, letter2.SideB):
		return "Both letters are equal!"
	case math.Min(letter1.SideA, letter1.SideB) >= math.Min(letter2.SideA, letter2.SideB) && math.Max(letter1.SideA, letter1.SideB) >= math.Max(letter2.SideA, letter2.SideB):
		return "You CAN put the second letter into the first"
	case math.Min(letter1.SideA, letter1.SideB) <= math.Min(letter2.SideA, letter2.SideB) && math.Max(letter1.SideA, letter1.SideB) <= math.Max(letter2.SideA, letter2.SideB):
		return "You CAN put the first letter into the second"
	default:
		return "You CAN'T put letters into each other"
	}
}

func printResult(result string) bool {
	if result != "You CAN'T put letters into each other" {
		fmt.Println(result)
		return true
	}
	fmt.Printf(result)
	return false
}

func askUserContinue() string {
	var answer string
	fmt.Println(`Would you like to continue the comparison? Enter "y" or "yes" to continue`)
	fmt.Scanf("%s", &answer)
	return answer
}

func analyzeAnswer(answer string) bool {
	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		return true
	}
	return false
}

func main() {
	askUserValues()
	result := compareLetters(letter1, letter2)
	printResult(result)
	answer := askUserContinue()
	confirm := analyzeAnswer(answer)
	if confirm {
		main()
	}
}
