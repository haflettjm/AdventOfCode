package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	directions := strings.Split(fileLines[0], "")
	return directions
}

func floorFinder(directions []string) int {
	count := 0
	for _, floor := range directions {
		if floor == "(" {
			count++
		} else {
			count--
		}
	}
	return count
}

func main() {
	filePath := "../input/input.txt"
	file := readInputFile(filePath)
	count := floorFinder(file)
	fmt.Println(count)
}
