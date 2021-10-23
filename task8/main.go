package main

import (
	"flag"
	"fmt"
	"strings"
)

const instructions = "---------------------------------------\n" +
	"Program has to launch with flags -l, -h \n" +
	"-l sets a lower value of range\n" +
	"-h sets a higher value of range\n" +
	"Appropriate values are higher than 1\n" +
	"---------------------------------------"

var(
	low, high int
)

func checkParams()bool {
	flag.IntVar(&low, "l", 0, "Lower value of range")
	flag.IntVar(&high, "h", 0, "Higher value of range")
	flag.Parse()
	return low > 0 && high > 0
}

func fibo() {
	f1, f2, f3 := 0, 1, 1
	var scope []string
	for i := 0; f3 < high; i++{
		f1 = f2
		f2 = f3
		f3 = f1 + f2
		if f3 >= low && f3 <= high {
			scope = append(scope, fmt.Sprint(f3))
		}
	}
	fmt.Printf(strings.Join(scope, ", "))
}

func main()  {
	if checkParams() {
		fibo()
		return
	}
	fmt.Println(instructions)
}

