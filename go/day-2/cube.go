package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var cubesInBag = map[string]int{
	"red":   12,
	"blue":  14,
	"green": 13,
}

func main() {
	var gamesMap = map[string]bool{}
	result := 0
	file := openFile()
	scanner := bufio.NewScanner(file)

	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		id := ""
		numberOfCurrentCubes := ""
		currentCubeColor := ""
		colonPassed := false

		for _, r := range line {
			if r == ':' {
				colonPassed = true
				gamesMap[id] = true
			}

			if unicode.IsDigit(r) && !colonPassed {
				id = id + string(r)
			}

			if colonPassed {
				if unicode.IsDigit(r) {
					numberOfCurrentCubes = numberOfCurrentCubes + string(r)
				} else if unicode.IsLetter(r) {
					currentCubeColor = currentCubeColor + string(r)
					if digit, err := strconv.Atoi(numberOfCurrentCubes); err == nil {
						switch currentCubeColor {
						case "red":
							if digit > cubesInBag["red"] {
								gamesMap[id] = false
							}
						case "green":
							if digit > cubesInBag["green"] {
								gamesMap[id] = false
							}
						case "blue":
							if digit > cubesInBag["blue"] {
								gamesMap[id] = false
							}
						}
					}
				} else if unicode.IsPunct(r) {
					currentCubeColor = ""
					numberOfCurrentCubes = ""
				}
			}
		}
	}

	for key, value := range gamesMap {
		if value {
			fmt.Println("", key)
			if digit, err := strconv.Atoi(key); err == nil {
				result += digit
			}
		}
	}
	fmt.Println("The result is:", result)
}

func openFile() *os.File {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return file
}
