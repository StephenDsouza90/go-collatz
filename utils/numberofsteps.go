package utils

// GetNumberOfSteps : A function that counts the number of steps it takes for any given positive number to reach 1.
// Based on a starting number,
// - if the number is even then "num / 2",
// - if the number is odd then "num * 3 + 1".
// :param userInput: any given positive number from N
// :return countSteps: the number of steps it takes to reach 1 for any given userInput
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
