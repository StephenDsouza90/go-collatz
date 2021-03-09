package utils

// GetNumberOfSteps : A function that counts the number of steps it takes for any given positive number to reach 1.
// userInput : Any given positive number from N
func GetNumberOfSteps(userInput int) int {
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
