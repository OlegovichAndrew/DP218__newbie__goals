package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var triangleList = make(map[string]float64)

func main() {
	askAndMake()
	printResult()
}

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(`Enter params of triangle in format: <name>, <side>, <side>, <side> need comma between as a delimiter: `)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func prepareInput(input string) (string, float64, float64, float64, error) {
	result := strings.Split(input, ",")
	if len(result) != 4 {
		return "", 0, 0, 0, errors.New("wrong number of arguments")
	}

	name := strings.ToLower(strings.TrimSpace(result[0]))
	a, err1 := strconv.ParseFloat(strings.TrimSpace(result[1]), 64)
	b, err2 := strconv.ParseFloat(strings.TrimSpace(result[2]), 64)
	c, err3 := strconv.ParseFloat(strings.TrimSpace(result[3]), 64)

	if err1 != nil || err2 != nil || err3 != nil {
		return name, 0, 0, 0, errors.New("wrong side value")
	}

	if a+b > c || a+c > b || b+c > a { // проверка на то, образуется ли треугольник с данными сторонами
		return name, a, b, c, nil
	}
	return name, a, b, c, errors.New(`can't build a triangle with inputted values'`)
}

func calculation(name string, a, b, c float64, err error) {
	if err != nil {
		fmt.Print("Error!: ", err)
		return
	}
	p := (a + b + c) / 2
	area := math.Sqrt(p * (p - a) * (p - b) * (p - c))

	if _, ok := triangleList[name]; !ok {
		triangleList[name] = area
		return
	}

	fmt.Print("Use only unique names for triangles")
}

func askUserContinue() {
	fmt.Println("\nEnter 'y' or 'yes' if you want to add one more triangle")
	var answer string
	fmt.Scan(&answer)
	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		askAndMake()
	}
}

func askAndMake() {
	input := userInput()
	calculation(prepareInput(input))
	askUserContinue()
}

func printResult() {
	names := make(map[float64][]string)
	var values []float64
	for k, v := range triangleList {
		names[v] = append(names[v], k)
	}
	for k := range names {
		values = append(values, k)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(values)))

	fmt.Println("===============Triangles list===============")
	for i, k := range values {
		for _, s := range names[k] {
			fmt.Printf("%d. [Triangle %s]: %.3f cm2\n", i+1, s, k)
		}
	}
}
