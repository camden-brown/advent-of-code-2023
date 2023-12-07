package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var Total int = 0
var CardNumberMap map[int]int = make(map[int]int)

func main() {
	file := openFile()
	scanner := bufio.NewScanner(file)

	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		countNumbers := false
		countWinningNumbers := false
		currentNumber := ""
		numberOfMatches := 0
		points := 1

		for i, r := range line {
			if r == ':' {
				countNumbers = true
			}

			if countNumbers {
				if unicode.IsDigit(r) {
					currentNumber += string(r)
				} else if r == ' ' && currentNumber != "" {
					CardNumberMap[convertToInt(currentNumber)] = 0
					currentNumber = ""
				} else if r == '|' {
					countNumbers = false
					countWinningNumbers = true
				}
			}

			if countWinningNumbers {
				if i+1 >= len(line) {
					currentNumber += string(r)
					convertedNumber := convertToInt(currentNumber)
					if _, ok := CardNumberMap[convertedNumber]; ok {
						numberOfMatches++
						CardNumberMap[convertedNumber]++
					} else {
						CardNumberMap[convertedNumber] = 0
					}
				} else if unicode.IsDigit(r) {
					currentNumber += string(r)
				} else if r == ' ' && currentNumber != "" {
					convertedNumber := convertToInt(currentNumber)
					if _, ok := CardNumberMap[convertedNumber]; ok {
						numberOfMatches++
						CardNumberMap[convertedNumber]++
					} else {
						CardNumberMap[convertedNumber] = 0
					}
					currentNumber = ""
				}
			}
		}

		if numberOfMatches > 0 {
			for i := 1; i < numberOfMatches; i++ {
				points *= 2
			}

			Total += points
		}

		countWinningNumbers = false
		countNumbers = false
		numberOfMatches = 0
		points = 1
		CardNumberMap = make(map[int]int)
	}

	fmt.Println("Winning Total:", Total)
}

func sumOfDigits(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

func openFile() *os.File {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func convertToInt(s string) int {
	digit, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return digit
}
