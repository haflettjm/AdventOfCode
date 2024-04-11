package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	return fileLines
}

func calcArea(num1 int, num2 int) int {
	return 2 * num1 * num2
}

func cleanDimensions(dimension string) []int {
	notNum := strings.Split(dimension, "x")
	var cleaned []int
	for _, num := range notNum {
		newNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
		}
		cleaned = append(cleaned, newNum)
	}
	return cleaned
}

func findLowest(num1 int, num2 int, num3 int) int {
	lowest := num1
	if num2 < num1 {
		lowest = num2
	}
	if num3 < num3 {
		lowest = num3
	}
	return lowest
}

func iterateOverData(data []string) {
	var sum int
	for _, line := range data {
		dimensions := cleanDimensions(line)
		x := dimensions[0]
		y := dimensions[1]
		z := dimensions[2]
		side1 := calcArea(x, y)
		side2 := calcArea(y, z)
		side3 := calcArea(x, z)
		lowest := findLowest(side1, side2, side3) / 2
		fmt.Printf("Dimensions are: %v, %v, %v \nThe lowest side is: %v\n", side1, side2, side3, lowest)
		areaNeeded := side1 + side2 + side3 + lowest
		fmt.Println(areaNeeded)
		sum += areaNeeded
	}
	fmt.Println(sum)
}

func main() {
	filePath := "../input/input.txt"
	file := readInputFile(filePath)
	iterateOverData(file)
}
