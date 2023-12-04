package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var prevLineSymbolsMap = map[int]rune{}
	currentNumber := ""
	result := 0
	file := openFile()
	scanner := bufio.NewScanner(file)

	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		for lineIndex, r := range line {
			if unicode.IsSymbol(r) || r == '*' || r == '#' {
				prevLineSymbolsMap[lineIndex] = r
				fmt.Println(unicode.IsPunct(r), string(r), currentNumber)
				if lineIndex > 0 {
					var prevIndexValue rune = rune(line[lineIndex-1])
					if unicode.IsDigit(rune(prevIndexValue)) {
						digit, err := strconv.Atoi(currentNumber)
						if err != nil {
							fmt.Println(err)
						}

						result += digit
					}
				}
				if lineIndex < len(line)-1 {
					var nextIndexValue rune = rune(line[lineIndex+1])
					if unicode.IsDigit(rune(nextIndexValue)) {
						slice := []rune(line)[:lineIndex]
						for _, r := range slice {
							if !unicode.IsDigit(r) {
								digit, err := strconv.Atoi(currentNumber)
								if err != nil {
									fmt.Println(err)
								}
								result += digit
								break
							}
							currentNumber = currentNumber + string(r)
						}
					}
				}

			}

			if unicode.IsDigit(r) {
				currentNumber = currentNumber + string(r)
			} else {
				currentNumber = ""
			}
		}
	}
	fmt.Println("The result is:", result)
}

func openFile() *os.File {
	file, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	return file
}
