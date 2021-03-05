package main

import (
	"fmt"
	"time"

	"github.com/StephenDsouza90/CollatzConjecture/utils"
)

// TODO: Fix timing

func main() {

	// Part C counts the number of steps to reach 1 for each number within a range of numbers.
	// The func OptimizedGetNumberOfStepsForEachM() uses an optimized algorithm for computation.

	var startValue int
	var endValue int

	fmt.Println("Enter start value: ")
	fmt.Scan(&startValue)

	fmt.Println("Enter end value: ")
	fmt.Scan(&endValue)

	start := time.Now()

	allSteps := utils.OptimizedGetNumberOfStepsForEachM(startValue, endValue)

	duration := time.Since(start)

	fmt.Printf("\nAll steps for each number within a range of numbers from %v to %v \n\n", startValue, endValue)
	fmt.Println(allSteps)

	fmt.Printf("\n--- Took %v seconds ---\n\n", duration.Seconds())

}