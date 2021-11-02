package main

import (
	"flag"
	"fmt"
	"strings"
)

var limit int

const instructions = "---------------------------------------------------------------\n" +
	"Program will print the natural numbers which squares are lesser than the limit\n" +
	"Set the limit with the flag -l\n" +
	"Example:flag -l 255 sets the limit = 255\n" +
	"---------------------------------------------------------------"

func adoptFlags() bool {
	flag.IntVar(&limit, "l", 0, "Upper limit for squares")
	flag.Parse()
	return limit > 0
}

func main() {
	if adoptFlags() {
		fmt.Println(printNaturalNumbers(limit))
		return
	}
	fmt.Println(instructions)
}

func printNaturalNumbers(limit int) string {
	var result []string
	for i := 1; i < limit; i++ {
		if i*i < limit {
			result = append(result, fmt.Sprint(i))
		}
	}
	return strings.Join(result, ",")
}
