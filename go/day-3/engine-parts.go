package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

//Array of Number struct
//Iterate through line
//If digit, create Number struct with value, start index, line
//Once no longer digit, add end index to Number struct
//Push Number struct to array

//Iterate through array of Number structs

type Number struct {
	value      string
	startIndex int
	endIndex   int
	line       int
	consumed   bool
}

type Symbol struct {
	index int
	line  int
}

var currentNumberStruct Number
var foundNumbers []Number
var foundSymbols []Symbol
var lineCounter int = 0
var currentLine, previousLine string
var result int = 0

func main() {
	file := openFile()
	scanner := bufio.NewScanner(file)

	defer file.Close()

	for scanner.Scan() {
		previousLine = currentLine
		currentLine = scanner.Text()

		for i, r := range currentLine {
			if isSpecialCharacter(r) {
				symbolStruct := Symbol{index: i, line: lineCounter}
				foundSymbols = append(foundSymbols, symbolStruct)
			}

			if unicode.IsDigit(r) {
				if currentNumberStruct != (Number{}) {
					currentNumberStruct.endIndex = i
					currentNumberStruct.value = currentNumberStruct.value + string(r)
				} else {
					currentNumberStruct = Number{value: string(r), startIndex: i, endIndex: i, line: lineCounter, consumed: false}
				}

				if i+1 >= len(currentLine) || !unicode.IsDigit(rune(currentLine[i+1])) {
					foundNumbers = append(foundNumbers, currentNumberStruct)
					currentNumberStruct = Number{}
				}
			}
		}
		lineCounter++
	}

	for i := range foundNumbers {
		val := &foundNumbers[i]
		for _, symbol := range foundSymbols {
			if symbol.line == val.line && val.consumed == false {
				if symbol.index == val.startIndex-1 || symbol.index == val.endIndex+1 {
					digit, err := strconv.Atoi(val.value)
					if err != nil {
						fmt.Println(err)
					}
					val.consumed = true
					result += digit
				}
			}

			if symbol.line == val.line-1 && val.consumed == false {
				if isInRange(symbol.index, val.startIndex-1, val.endIndex+1) {
					digit, err := strconv.Atoi(val.value)
					if err != nil {
						fmt.Println(err)
					}
					val.consumed = true
					result += digit
				}
			}

			if symbol.line == val.line+1 && val.consumed == false {
				if isInRange(symbol.index, val.startIndex-1, val.endIndex+1) {
					digit, err := strconv.Atoi(val.value)
					if err != nil {
						fmt.Println(err)
					}

					val.consumed = true
					result += digit
				}
			}
		}
	}
	// fmt.Println("The result is:", result, foundNumbers)
	// fmt.Println("The result is:", result, foundSymbols)
	fmt.Println("The result is:", result)
}

func isInRange(n, start, end int) bool {
	return n >= start && n <= end
}

func openFile() *os.File {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func isSpecialCharacter(r rune) bool {
	specialCharacters := []rune{'*', '+', '=', '&', '%', '$', '@', '#', '-', '/', '&'}
	for _, c := range specialCharacters {
		if r == c {
			return true
		}
	}
	return false
}
