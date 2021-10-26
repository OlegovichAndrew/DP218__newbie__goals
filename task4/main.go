package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const instructions = `Program starts with command line.
First mode start with 2 parameters: <path to file> <string to count>;
Second mode start with 3 parameters: <path to file> <string to find> <string to write>`

func main() {
	if len(os.Args) == 3 {
		readAndCount(os.Args[1], os.Args[2])
	} else if len(os.Args) == 4 {
		rewriteString(os.Args[1], os.Args[2], os.Args[3])
	} else {
		fmt.Printf("%s\n", instructions)
	}
}

func readAndCount(file, stringToFind string) {
	var counter int
	fileData, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer fileData.Close()

	read := bufio.NewReader(fileData)

	for {
		line, err := read.ReadString('\n')
		counter += strings.Count(line, stringToFind)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error was happened: ", err)
			break
		}
	}
	fmt.Printf("The text includs the desired string %d times\n", counter)
	return
}

func rewriteString(file, stringToFind, stringToWrite string) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Errorf("error happened %q", err)
		return
	}

	newContent := string(content)
	newContent = strings.Replace(newContent, stringToFind, stringToWrite, -1)
	temp, err := os.Create("temp.txt")
	if err != nil {
		fmt.Errorf("error happened %q", err)
		return
	}

	_, err = temp.WriteString(newContent)
	if err != nil {
		fmt.Errorf("error happened %q", err)
		return
	}

	err = os.Remove("file.txt")
	if err != nil {
		fmt.Errorf("error happened %q", err)
		return
	}

	err = os.Rename("temp.txt", "file.txt")
	if err != nil {
		fmt.Errorf("error happened %q", err)
		return
	}
	fmt.Printf("String was edited\n")
}
