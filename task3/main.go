package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const instructions = `---------------------------------------------------------------------
You should write parameters in the next order: <name>, <side>, <side>, <side>
Then program will calculate the result for the triangle.
You can add triangles as much as you want one after another.
In progress program will ask your confirmation.
Finally it shows the list of triangles in a sorted order`

var triangleList = make(map[string]float64)

func main() {
	input := userInput(os.Stdin)
	name, area, err := formulaHeron(prepareInput(input))
	triangleList = addToSlice(name, area, err)
	if askUserContinue() {
		main()
		return
	}
	printResult(triangleList)
}

func userInput(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	fmt.Print(`Enter parameters of triangle in format: <name>, <side>, <side>, <side> need comma between as a delimiter: `)
	var input string
	scanner.Scan()
	input = scanner.Text()
	if input == "" {
		fmt.Println(instructions)
		return ""
	}
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

	if a+b > c && a+c > b && b+c > a { // проверка на то, образуется ли треугольник с данными сторонами
		return name, a, b, c, nil
	}
	return "", a, b, c, errors.New(`can't build a triangle with inputted values'`)
}

func formulaHeron(name string, a, b, c float64, err error) (string, float64, error) {
	if err != nil {
		fmt.Print("Error!: ", err)
		return "", 0, err
	}
	p := (a + b + c) / 2
	area := math.Round((math.Sqrt(p*(p-a)*(p-b)*(p-c)))*100) / 100
	return name, area, nil
}

func addToSlice(name string, area float64, err error) map[string]float64 {
	if err != nil {
		return triangleList
	}
	if _, ok := triangleList[name]; ok {
		fmt.Print("Use only unique names for triangles")
		return triangleList
	}
	triangleList[name] = area
	return triangleList
}

func askUserContinue() bool {
	fmt.Println("\nEnter 'y' or 'yes' if you want to add one more triangle")
	var answer string
	fmt.Scan(&answer)
	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		return true
	}
	return false
}

func printResult(triangleList map[string]float64) {
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
			fmt.Printf("%d. [Triangle %s]: %.2f cm2\n", i+1, s, k)
		}
	}
}
