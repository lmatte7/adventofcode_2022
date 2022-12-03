package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	priorityTotal := 0

	for scanner.Scan() {
		ruckList := scanner.Text()

		rowPriority := 0

		firstHalf := ruckList[0 : len(ruckList)/2]
		secondHalf := ruckList[(len(ruckList) / 2):]

		commonLetter := ""

		for _, letter := range strings.Split(firstHalf, "") {
			if strings.Contains(secondHalf, letter) {
				commonLetter = letter
				break
			}
		}

		// Convert the common letter to ASCII and subtract to get the score values
		priorityRune := []rune(commonLetter)
		if unicode.IsUpper(priorityRune[0]) {
			rowPriority += int(priorityRune[0]) - 38
		} else {
			rowPriority += int(priorityRune[0]) - 96
		}

		fmt.Printf("First Half: %v, Second Half: %v, Common Letter: %v, Priority: %v\n", firstHalf, secondHalf, commonLetter, rowPriority)

		priorityTotal += rowPriority
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return priorityTotal
}
