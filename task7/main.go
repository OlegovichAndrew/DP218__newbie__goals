package main

import (
	"flag"
	"fmt"
	"strings"
)

var limit int

const instructions = "---------------------------------------------------------------\n" +
	"Program will print squares of natural numbers until the limit\n" +
	"Set the limit with the flag -l\n" +
	"Example:flag -l 255 sets the limit = 255\n" +
	"---------------------------------------------------------------"

func adoptFlags() bool {
	flag.IntVar(&limit, "l", 0, "Upper limit for squares")
	flag.Parse()
	isNatural := limit > 0
	return isNatural
}

func main()  {
	if adoptFlags(){
		printNaturalNumbers()
	} else {
		fmt.Println(instructions)
	}

}

func printNaturalNumbers() {
	var result []string
	for i := 1; i < limit; i ++ {
		if i * i < limit {
			result = append(result, fmt.Sprint(i))
		}
	}
	fmt.Printf(strings.Join(result, ","))
}




