package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const instruction = `There is the program which analyzes tickets, and writes if the one is LUCKY.
You have to enter the path to file where situated one of the keys: "moscow" or "piter"`

var ticket, ticket2 = "044530", "356536"

func main() {

	file := os.Args[1]
	parameter := readFile(file)

	switch parameter {
	case "moscow":
		luck := moscowAlgorythm(ticket)
		if luck {
			fmt.Printf("The ticket %s is lucky with Moscow counting\n", ticket)
		} else {
			fmt.Printf("The ticket %s is NOT lucky with Moscow counting\n", ticket)
		}
	case "piter":
		luck := piterAlgorythm(ticket)
		if luck {
			fmt.Printf("The ticket %s is lucky with Piter counting\n", ticket)
		} else {
			fmt.Printf("The ticket %s is not lucky with Piter counting\n", ticket)
		}
	default:
		fmt.Println("Can't run the program with this parameter")
	}
}

func readFile(input string) string {
	data, err := os.ReadFile("moscow.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	parameter := string(data)
	return parameter
}

func piterAlgorythm(ticket string) bool {
	var odd, even int64
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
		return true
	}
	return false
}

func moscowAlgorythm(ticket string) bool {
	var left, right int64

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
		return true
	}
	return false
}
