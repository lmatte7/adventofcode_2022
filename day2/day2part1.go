package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	points := parseInput("input.txt")

	fmt.Printf("Total Points: %d\n", points)

}

func parseInput(filename string) int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalPoints := 0

	for scanner.Scan() {
		playRunes := []rune(scanner.Text())

		humanChoice := int(playRunes[2]) - 23

		fmt.Printf("Choices: %d, %d\n", int(playRunes[0]), humanChoice)

		// Add choice score by subtracting from ASCII value
		totalPoints += humanChoice - 64
		if int(playRunes[0]) == humanChoice {
			totalPoints += 3
		} else {
			if humanChoice == 65 {
				if int(playRunes[0]) == 67 {
					totalPoints += 6
				}
			} else if humanChoice == 66 {
				if int(playRunes[0]) == 65 {
					totalPoints += 6
				}
			} else if humanChoice == 67 {
				if int(playRunes[0]) == 66 {
					totalPoints += 6
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalPoints
}
