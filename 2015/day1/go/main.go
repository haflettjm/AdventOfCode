package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputFile(path string) []string {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		fmt.Println(line)
	}

	return fileLines
}

func main() {
	filePath := "../input/input.txt"
	fmt.Println(readInputFile(filePath))

}
