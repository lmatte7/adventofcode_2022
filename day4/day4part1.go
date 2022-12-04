package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	overlaps := parseInput("input.txt")

	fmt.Printf("Overlapping Assigments: %d\n", overlaps)

}

func parseInput(filename string) int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	overlappingAssignments := 0

	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")

		firstAssignment := strings.Split(assignments[0], "-")
		secondAssignment := strings.Split(assignments[1], "-")

		// Loop through assignments to make unique strings
		firstAssignmentStart, err := strconv.Atoi(firstAssignment[0])
		if err != nil {
			panic(err)
		}
		firstAssignmentEnd, err := strconv.Atoi(firstAssignment[1])
		if err != nil {
			panic(err)
		}

		secondAssignmentStart, err := strconv.Atoi(secondAssignment[0])
		if err != nil {
			panic(err)
		}
		secondAssignmentEnd, err := strconv.Atoi(secondAssignment[1])
		if err != nil {
			panic(err)
		}

		if firstAssignmentStart <= secondAssignmentStart && firstAssignmentEnd >= secondAssignmentEnd {
			overlappingAssignments++
			fmt.Printf("First %d/%d, Second %d/%d\n\n", firstAssignmentStart, firstAssignmentEnd, secondAssignmentStart, secondAssignmentEnd)
		} else if secondAssignmentStart <= firstAssignmentStart && secondAssignmentEnd >= firstAssignmentEnd {
			overlappingAssignments++
			fmt.Printf("First %d/%d, Second %d/%d\n\n", firstAssignmentStart, firstAssignmentEnd, secondAssignmentStart, secondAssignmentEnd)
		}

	}

	return overlappingAssignments
}
