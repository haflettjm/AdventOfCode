package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputFile(path string)string{
	scanner := bufio.NewScanner(r Reader)
	readFile, err := os.Open(path)
	if err != nil{
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines{
		fmt.Println(line)
	}

	fmt.Println(fileLines)
}

func main() {
	filePath := "../input/input.txt"
	floor := 0
	fmt.Println(floor)


}
