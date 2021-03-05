package main

import (
	"fmt"

	"github.com/StephenDsouza90/CollatzConjecture/utils"
)

func main() {

	// Part A counts the number of steps it takes for any given positive number to reach 1.
	// The steps follow a sequence defined by "Collatz conjecture" theory (See README).
	// For explanation on the GetNumberOfSteps func
	// see github.com/StephenDsouza90/CollatzConjecture/utils/numberofsteps.go

	var userInput int

	fmt.Println("Enter value for m: ")
	fmt.Scanln(&userInput)

	steps := utils.GetNumberOfSteps(userInput)
	fmt.Printf("It takes %v steps for the number %v to reach the number 1.\n", steps, userInput)

}
