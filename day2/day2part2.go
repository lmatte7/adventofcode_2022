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
		if humanChoice == 65 {
			// totalPoints += 3
			if int(playRunes[0]) == 65 {
				totalPoints += 3
			}
			if int(playRunes[0]) == 66 {
				totalPoints += 1
			}
			if int(playRunes[0]) == 67 {
				totalPoints += 2
			}
		} else if humanChoice == 66 {
			totalPoints += 3
			// We need to draw so add the same choice
			totalPoints += int(playRunes[0]) - 64
		} else if humanChoice == 67 {
			totalPoints += 6

			if int(playRunes[0]) == 65 {
				totalPoints += 2
			}
			if int(playRunes[0]) == 66 {
				totalPoints += 3
			}
			if int(playRunes[0]) == 67 {
				totalPoints += 1
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalPoints
}
