package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	caloriesArray := parseInput("input.txt")

	sort.Ints(caloriesArray)

	cLen := len(caloriesArray)
	fmt.Println(caloriesArray)

	topThreeCalories := caloriesArray[cLen-1] + caloriesArray[cLen-2] + caloriesArray[cLen-3]

	fmt.Printf("Top Three Elfs: %v\n: ", topThreeCalories)
}

func parseInput(filename string) []int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	values := make([]int, 0)

	scanner := bufio.NewScanner(file)
	elfCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			values = append(values, elfCalories)
			elfCalories = 0
		} else {
			intValue, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			elfCalories += intValue
		}

	}
	values = append(values, elfCalories)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values
}
