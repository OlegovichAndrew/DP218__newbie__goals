package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const instruction = `There is the program which analyzes tickets, and writes if the one is LUCKY.
You have to enter the path to file where situated one of the keys: "moscow" or "piter"`

func main() {
	input := readInput()
	parameter := readFile(input)

	switch parameter {
	case "moscow":
		moscowAlgorythm()
	case "piter":
		piterAlgorythm()
	default:
		fmt.Println("Can't run the program with this parameter")
	}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(`enter the file path`)
	scanner.Scan()
	if scanner.Text() == "" {
		fmt.Println(instruction)
		readInput()
	}
	input := scanner.Text()
	return input
}

func readFile(input string) string {
	data, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	parameter := string(data)
	return parameter
}

func piterAlgorythm() {
	var odd, even int64
	ticket := "044530"
	split := strings.Split(ticket, "")

	for _, v := range split {
		t, _ := strconv.ParseInt(v, 0, 64)
		if t%2 == 0 {
			even += t
		} else {
			odd += t
		}
	}

	if odd == even {
		fmt.Printf("The ticket %s is lucky with Piter counting\n", ticket)
		return
	}
	fmt.Printf("The ticket %s is not lucky with Piter counting\n", ticket)
}

func moscowAlgorythm() {
	var left, right int64
	ticket := "356536"
	split := strings.Split(ticket, "")

	for i, v := range split {
		if i < 3 {
			t, _ := strconv.ParseInt(v, 0, 64)
			left += t
		} else {
			t, _ := strconv.ParseInt(v, 0, 64)
			right += t
		}
	}

	if left == right {
		fmt.Printf("The ticket %s is lucky with Moscow counting\n", ticket)
		return
	}
	fmt.Printf("The ticket %s is NOT lucky with Moscow counting\n", ticket)
}
