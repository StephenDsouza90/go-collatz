package utils

// GetNumberOfSteps : A function that counts the number of steps it takes for any given positive number to reach 1.
func GetNumberOfSteps(userInput int) int {
	// userInput : any given positive number from N

	countSteps := 0

	// Based on a starting number,
	// - if the number is even then "num / 2",
	// - if the number is odd then "num * 3 + 1".
	for userInput != 1 {
		if userInput%2 == 0 {
			userInput = userInput / 2
			countSteps++
		} else {
			userInput = userInput*3 + 1
			countSteps++
		}
	}

	// countSteps : the number of steps it takes to reach 1 for any given userInput
	return countSteps
}
