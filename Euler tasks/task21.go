package main

import "fmt"

var sum int

func d(number int) int {
	var value int
	for i := 1; i < number; i++ {
		if number%i == 0 {
			value += i
		}
	}
	return value
}

func find() {
	for n := 1; n < 10000; n++ { // 220
		b := d(n) // 284
		a := d(b)
		if n == a && n != b {
			sum += n
			fmt.Printf("The d(%d) is: %d. The d(%d) is: %d. They are equal\n", n, b, b, a)
		}
	}
	fmt.Printf("The sum is: %d\n", sum)
}

func main() {
	find()
}
