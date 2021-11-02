package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const instructions = `Program starts with command line.
The first mode starts with 2 parameters: <path to file> <string to count>;
The second mode starts with 3 parameters: <path to file> <string to find> <string to write>`

func main() {
	if len(os.Args) == 3 {
		count, err := readAndCount(os.Args[1], os.Args[2])
		if err != nil {
			fmt.Printf("Error was happened: %v", err)
		}
		fmt.Printf("The text includs the desired string %d times\n", count)
	} else if len(os.Args) == 4 {
		err := rewriteString(os.Args[1], os.Args[2], os.Args[3])
		if err != nil {
			fmt.Printf("Error happened: %v", err)
		}
	} else {
		fmt.Printf("%s\n", instructions)
	}
}

func readAndCount(file, stringToFind string) (int, error) {
	var counter int
	fileData, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer fileData.Close()

	read := bufio.NewReader(fileData)

	for {
		line, err := read.ReadString('\n')
		counter += strings.Count(line, stringToFind)

		if err == io.EOF {
			return counter, nil
		} else if err != nil {
			return 0, err
		}
	}
}

func rewriteString(file, stringToFind, stringToWrite string) error {
	fmt.Println("Search started...")
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	newContent := string(content)
	newContent = strings.Replace(newContent, stringToFind, stringToWrite, -1)
	temp, err := os.Create("temp.txt")

	if err != nil {
		return err
	}
	_, err = temp.WriteString(newContent)
	if err != nil {
		return err
	}
	err = os.Remove("file.txt")
	if err != nil {
		return err
	}
	err = os.Rename("temp.txt", "file.txt")
	if err != nil {
		return err
	}
	fmt.Printf("Search and replace finished\n")
	return nil
}
