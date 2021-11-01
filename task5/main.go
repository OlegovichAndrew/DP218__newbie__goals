package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	tillTwenty = map[string]string{
		"0":  "",
		"1":  "один ",
		"2":  "два ",
		"3":  "три ",
		"4":  "четыре ",
		"5":  "пять ",
		"6":  "шесть ",
		"7":  "семь ",
		"8":  "восемь ",
		"9":  "девять ",
		"10": "десять ",
		"11": "одиннадцать ",
		"12": "двенадцать ",
		"13": "тринадцать ",
		"14": "четырнадцать ",
		"15": "пятнадцать ",
		"16": "шестнадцать ",
		"17": "семнадцать ",
		"18": "восемнадцать ",
		"19": "девятнадцать ",
	}

	tillHundred = map[string]string{
		"2": "двадцать ",
		"3": "тридцать ",
		"4": "сорок ",
		"5": "пятьдесят ",
		"6": "шестьдесят ",
		"7": "семьдесят ",
		"8": "восемьдесят ",
		"9": "девяносто ",
	}

	hundreds = map[string]string{
		"0": "",
		"1": "сто ",
		"2": "двести ",
		"3": "триста ",
		"4": "четыреста ",
		"5": "пятьсот ",
		"6": "шестьсот ",
		"7": "семьсот ",
		"8": "восемьсот ",
		"9": "девятьсот ",
	}
)

func userInput(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		input := scanner.Text()
		return input, nil
	}
	return "", scanner.Err()
}

// The next function devides input into pieces, depends on it length. It should be comfortable to work next.

func inputOperations(input string) ([]string, error) {
	var forRead []string
	_, err := strconv.Atoi(input)
	if err != nil {
		return forRead, errors.New("use only numbers in input")
	}
	// "empty" is a string to fill the slice till 5 elements
	empty := "000"

	switch {
	case len(input) >= 1 && len(input) <= 3: // 0-999 hundred
		piece1 := input
		forRead = append(forRead, empty, empty, empty, empty, piece1)
		return forRead, nil
	case len(input) >= 4 && len(input) <= 6: // 1.000 - 999.999 thousand
		piece1 := input[len(input)-3:]
		piece2 := input[:len(input)-3]
		forRead = append(forRead, empty, empty, empty, piece2, piece1)
		return forRead, nil
	case len(input) >= 7 && len(input) <= 9: // 1.000.000 - 999.999.999 million
		piece1 := input[len(input)-3:]
		piece2 := input[len(input)-6 : len(input)-3]
		piece3 := input[:len(input)-6]
		forRead = append(forRead, empty, empty, piece3, piece2, piece1)
		return forRead, nil
	case len(input) >= 10 && len(input) <= 12: // 1.000.000.000 - 999.999.999.999 billion
		piece1 := input[len(input)-3:]
		piece2 := input[len(input)-6 : len(input)-3]
		piece3 := input[len(input)-9 : len(input)-6]
		piece4 := input[:len(input)-9]
		forRead = append(forRead, empty, piece4, piece3, piece2, piece1)
		return forRead, nil
	case len(input) >= 13 && len(input) <= 15: // 1.000.000.000.000 - 999.999.999.999.999 trillion
		piece1 := input[len(input)-3:]
		piece2 := input[len(input)-6 : len(input)-3]
		piece3 := input[len(input)-9 : len(input)-6]
		piece4 := input[len(input)-12 : len(input)-9]
		piece5 := input[:len(input)-12]
		forRead = append(forRead, piece5, piece4, piece3, piece2, piece1)
		return forRead, nil
	default:
		return forRead, errors.New("cannot handle your number")
	}
}

// The next function adds zeros if length lesser than 3
func addToThree(piece string) string {
	chunk := piece
	if len(chunk) == 2 {
		chunk = "0" + piece
	}
	if len(chunk) == 1 {
		chunk = "00" + piece
	}
	return chunk
}

func logic(piece string) string {
	var toWrite string
	chunk := piece
	chunk = addToThree(chunk)

	// writes down quantity of hundreds
	toWrite += hundreds[string(chunk[0])]
	// trick to separate teens
	if chunk[1] == '1' {
		toWrite += tillTwenty[chunk[1:]]
		return toWrite
	}
	// fills the string with other numbers
	toWrite += tillHundred[string(chunk[1])] + tillTwenty[string(chunk[2])]
	return toWrite
}

func logicThousand(piece string) string {
	tillTwenty["1"] = "одна "
	tillTwenty["2"] = "две "
	result := logic(piece)
	tillTwenty["1"] = "один "
	tillTwenty["2"] = "два "
	return result
}

func addEnding(piece string) string {
	chunk := piece
	chunk = addToThree(chunk)

	switch {
	case string(chunk[2]) == "1" && string(chunk[1]) != "1":
		return " "
	case string(chunk[2]) == "2" && string(chunk[1]) != "1" || string(chunk[2]) == "3" && string(chunk[1]) != "1" || string(chunk[2]) == "4" && string(chunk[1]) != "1":
		return "а "
	default:
		return "ов "
	}
}

func addEndingThousand(piece string) string {
	chunk := piece
	chunk = addToThree(chunk)

	switch {
	case string(chunk[2]) == "1" && string(chunk[1]) != "1":
		return "а "
	case string(chunk[2]) == "2" && string(chunk[1]) != "1" || string(chunk[2]) == "3" && string(chunk[1]) != "1" || string(chunk[2]) == "4" && string(chunk[1]) != "1":
		return "и "
	default:
		return " "
	}
}

func completePieces(input []string) string {
	var result string
	var isZero int

	for _, v := range input {
		tmp, _ := strconv.Atoi(v)
		isZero += tmp
	}
	// verification if the input is 0
	if isZero == 0 {
		result = "ноль"
		return result
	}
	if input[0] != "000" && input[0] != "00" && input[0] != "0" {
		result += logic(input[0]) + "триллион" + addEnding(input[0])
	}
	if input[1] != "000" && input[1] != "00" && input[1] != "0" {
		result += logic(input[1]) + "миллиард" + addEnding(input[1])
	}
	if input[2] != "000" && input[2] != "00" && input[2] != "0" {
		result += logic(input[2]) + "миллион" + addEnding(input[2])
	}
	if input[3] != "000" && input[3] != "00" && input[3] != "0" {
		result += logicThousand(input[3]) + "тысяч" + addEndingThousand(input[3])
	}
	if input[4] != "000" {
		result += logic(input[4])
	}
	return result
}

func main() {
	str, err := userInput(os.Stdin)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	arr, err := inputOperations(str)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(completePieces(arr))
}
