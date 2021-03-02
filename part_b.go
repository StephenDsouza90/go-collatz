package main

import (
	"fmt"
	"time"
)

// TODO: Use import for getNumberOfSteps from part_a.go
// TODO: Find a better way for range

func main() {

	var startValue int
	var endValue int

	fmt.Println("Enter start value: ")
	fmt.Scan(&startValue)

	fmt.Println("Enter end value: ")
	fmt.Scan(&endValue)

	start := time.Now()

	allSteps := getNumberOfStepsForEachM(startValue, endValue)

	duration := time.Since(start)

	fmt.Printf("\nAll steps for each number within a range of numbers from %v to %v \n\n", startValue, endValue)
	fmt.Println(allSteps, "\n")

	fmt.Printf("--- Took %v seconds ---\n", duration.Seconds())

}

func getNumberOfStepsForEachM(startValue, endValue int) map[int]int {

	/* For each number num within a range of numbers between startValue and endValue,
	this function returns the number of steps.

	The number and its corresponding steps are stored in a dict,
	where key is number num and steps to compute the number is the value.

	:param startValue: starting number of the range
	:param endValue: ending number of the range
	:return allSteps: dict with m as the key and it's corresponding steps as the value */

	allSteps := make(map[int]int)

	for num := startValue; num <= endValue; num++ {
		steps := getNumberOfSteps(num)
		allSteps[num] = steps
	}

	return allSteps

}

func getNumberOfSteps(num int) int {

	countSteps := 0

	for num != 1 {

		if num%2 == 0 {
			num = num / 2
			countSteps++
		} else {
			num = num*3 + 1
			countSteps++
		}
	}

	return countSteps

}
