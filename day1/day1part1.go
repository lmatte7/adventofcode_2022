package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	caloriesArray := parseInput("input.txt")

	maxCalories := 0
	elfWitTheMost := -1
	for i, calories := range caloriesArray {
		if calories > maxCalories {
			elfWitTheMost = i
			maxCalories = calories
		}
	}

	fmt.Printf("Calories Array: %v\n", caloriesArray)
	fmt.Printf("Calories held by Elf Wit Da Most: %v\n", caloriesArray[elfWitTheMost])
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values
}
