package main

import (
	"fmt"
	"time"

	"github.com/StephenDsouza90/CollatzConjecture/utils"
)

// Part B counts the number of steps to reach 1 for each number within a range of numbers.
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
	fmt.Println(allSteps)

	fmt.Printf("\n--- Took %v seconds ---\n\n", duration.Seconds())

}

// getNumberOfStepsForEachM : For each number num within a range of numbers between startValue and endValue,
// this function returns the number of steps.
// The number and its corresponding steps are stored in a dict,
// where key is number num and steps to compute the number is the value.
// :param startValue: starting number of the range
// :param endValue: ending number of the range
// :return allSteps: dict with m as the key and it's corresponding steps as the value
func getNumberOfStepsForEachM(startValue, endValue int) map[int]int {

	allSteps := make(map[int]int)

	for num := startValue; num <= endValue; num++ {
		steps := utils.GetNumberOfSteps(num)
		allSteps[num] = steps
	}

	return allSteps

}
