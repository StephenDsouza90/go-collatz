package main

import (
	"fmt"
)

func main() {

	var userInput int

	fmt.Println("Enter value for m: ")
	fmt.Scanln(&userInput)

	steps := getNumberOfSteps(userInput)
	fmt.Printf("It takes %v steps for the number %v to reach the number 1.\n", steps, userInput)

}

func getNumberOfSteps(userInput int) int {
	/* A function that counts the number of steps it takes for
		any given positive number to reach 1.

		The steps follow a sequence defined by
		"Collatz conjecture" theory (See README).

		Based on a starting number,
		if the number is even then "num / 2",
	    if the number is odd then "num * 3 + 1".

		:param userInput: any given positive number from N
		:return countSteps: the number of steps it takes to reach 1 for any given userInput */

	countSteps := 0

	for userInput != 1 {

		if userInput%2 == 0 {
			userInput = userInput / 2
			countSteps++
		} else {
			userInput = userInput*3 + 1
			countSteps++
		}
	}

	return countSteps

}
