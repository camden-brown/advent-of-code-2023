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

var digitWords = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
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
		var currentWord string

		for i, r := range line {
			if unicode.IsDigit(r) {
				arr = append(arr, string(r))
				currentWord = ""
			} else {
				currentWord += string(r)
				if digit, exists := digitWords[currentWord]; exists {
					arr = append(arr, digit)
					// Check if the next letter has the prefix of the currentWords last letter
					if strings.HasPrefix(line[i+1], digitWords[]) {
						currentWord = ""
					} else {
						currentWord = ""
					}
				}
			}
			fmt.Println(currentWord)
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
		fmt.Println(sum, line, arr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The result is:", total)
}
