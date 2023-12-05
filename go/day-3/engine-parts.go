package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var currentNumber string = ""
var currentLine, previousLine string
var result int = 0
var numsAdd []int
var currentNumberIndexes []int

func main() {
	// var previousLineNumbersByIndex = map[int][]int{}
	file := openFile()
	scanner := bufio.NewScanner(file)

	defer file.Close()

	for scanner.Scan() {
		previousLine = currentLine
		currentLine = scanner.Text()

		for lineIndex, r := range currentLine {
			if IsSymbol(r) {
				symbolAdjacent(lineIndex)
			}

			if unicode.IsDigit(r) {
				currentNumber = currentNumber + string(r)
				currentNumberIndexes = append(currentNumberIndexes, lineIndex)
			} else {
				symbolDiagonalToDigit()
				currentNumber = ""
			}
		}
		currentNumberIndexes = make([]int, 0)
	}
	fmt.Println("The result is:", result, numsAdd)
}

func openFile() *os.File {
	file, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func IsSymbol(r rune) bool {
	return unicode.IsSymbol(r) || r == '*' || r == '#'
}

func symbolDiagonalToDigit() {
	if currentNumber != "" && previousLine != "" {
		for _, val := range currentNumberIndexes {
			digit, err := strconv.Atoi(currentNumber)
			if err != nil {
				fmt.Println(err)
			}

			if IsSymbol(rune(previousLine[val])) || (val > 0 && IsSymbol(rune(previousLine[val-1]))) || IsSymbol(rune(previousLine[val+1])) {

				numsAdd = append(numsAdd, digit)

				result += digit
				break
			}
		}
	}
}

func symbolAdjacent(lineIndex int) {
	if lineIndex > 0 {
		var prevIndexValue rune = rune(currentLine[lineIndex-1])
		if unicode.IsDigit(rune(prevIndexValue)) {
			digit, err := strconv.Atoi(currentNumber)
			if err != nil {
				fmt.Println(err)
			}
			numsAdd = append(numsAdd, digit)
			result += digit
		}
	}

	if lineIndex < len(currentLine)-1 {
		var nextIndexValue rune = rune(currentLine[lineIndex+1])
		if unicode.IsDigit(rune(nextIndexValue)) {
			slice := []rune(currentLine)[:lineIndex]
			for _, r := range slice {
				if !unicode.IsDigit(r) {
					digit, err := strconv.Atoi(currentNumber)
					if err != nil {
						fmt.Println(err)
					}
					numsAdd = append(numsAdd, digit)
					result += digit
					break
				}
				currentNumber = currentNumber + string(r)
			}
		}
	}
}
