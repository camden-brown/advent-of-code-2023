package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var DigitWords = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	var total int = 0
	file, err := os.Open("puzzle.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		var arr []string
		var sum int = 0

		for _, word := range DigitWords {
			var wordToDigit string = convertDigitWordToStringNumber(word)
			line = strings.Replace(line, word, wordToDigit, -1)
		}
		fmt.Println(line)

		for _, r := range line {
			if unicode.IsDigit(r) {
				arr = append(arr, string(r))
			}
		}

		var arrLength int = len(arr)
		var combinedDigitString string = ""

		if arrLength == 1 {
			combinedDigitString = arr[0] + arr[0]
		} else {
			combinedDigitString = arr[0] + arr[arrLength-1]
		}

		if digit, err := strconv.Atoi(combinedDigitString); err == nil {
			sum = digit
		}
		total += sum
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The result is:", total)
}

func convertDigitWordToStringNumber(arg string) string {
	var result string
	switch arg {
	case "one":
		result = "1"
	case "two":
		result = "2"
	case "three":
		result = "3"
	case "four":
		result = "4"
	case "five":
		result = "5"
	case "six":
		result = "6"
	case "seven":
		result = "7"
	case "eight":
		result = "8"
	case "nine":
		result = "9"
	}
	return result
}
