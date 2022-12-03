package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {

	total := parseInput("input.txt")

	fmt.Printf("Total Points: %d\n", total)

}

func parseInput(filename string) int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rowCount := 0
	commonRunes := make(map[rune]int, 0)
	rowPriority := 0
	priorityRune := '1'
	priorityTotal := 0

	for scanner.Scan() {
		ruckList := scanner.Text()

		rowRunes := make(map[rune]int, 0)
		for _, letter := range ruckList {
			if rowRunes[letter] == 0 {
				commonRunes[letter]++
				rowRunes[letter]++
			}
		}

		rowCount++

		if rowCount == 3 {

			for pRune, frequency := range commonRunes {
				if frequency == 3 {
					priorityRune = pRune
				}
			}

			if unicode.IsUpper(priorityRune) {
				rowPriority = int(priorityRune) - 38
				fmt.Printf("Letter Prioty: %v ", int(priorityRune)-38)
			} else {
				rowPriority = int(priorityRune) - 96
				fmt.Printf("Letter Prioty: %v ", int(priorityRune)-96)
			}
			fmt.Printf("Priority Letter: %v,\n", string(priorityRune))

			priorityTotal += rowPriority
			commonRunes = make(map[rune]int, 0)
			rowCount = 0
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return priorityTotal
}
