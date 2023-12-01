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
			fmt.Println("sum:", digit)
		}
		total += sum
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The result is:", total)
}
