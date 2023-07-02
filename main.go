package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/day1.input")

	if err != nil {
		fmt.Println("error reading file: ", err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
 
	maxCalories := math.MinInt
	currentCalories := 0
	for fileScanner.Scan() {
		
		line := fileScanner.Text();

		if len(strings.TrimSpace(line)) == 0 {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		} else {
			calories, err := strconv.Atoi(line) 
			if err != nil {
				fmt.Println("line is not i32: ", err)
			}
			currentCalories += calories
		}
	}
 
	file.Close()
	fmt.Println("max calories: ", maxCalories)
}