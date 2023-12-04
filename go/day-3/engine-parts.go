package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var gamesMap = map[string]bool{}
	result := 0
	file := openFile()
	scanner := bufio.NewScanner(file)

	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		nextLine()

		for _, r := range line {

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

{
	[lineId]: {
		[index]: number
	}
}
