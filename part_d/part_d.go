package main

import (
	"fmt"

	"github.com/StephenDsouza90/CollatzConjecture/utils"

	"github.com/Arafatk/glot"
)

func main() {

	var startValue int
	var endValue int

	fmt.Println("Enter start value: ")
	fmt.Scan(&startValue)

	fmt.Println("Enter end value: ")
	fmt.Scan(&endValue)

	allSteps := utils.OptimizedGetNumberOfStepsForEachM(startValue, endValue)

	listX, listY := getAxis(allSteps)

	plotGraph(listX, listY)

	fmt.Printf("\nOpen Collatz.png to view plotted diagram.\n")

}

func getAxis(allSteps map[int]int) ([]float64, []float64) {

	listX := []float64{}
	listY := []float64{}

	for key, value := range allSteps {
		listX = append(listX, float64(key))
		listY = append(listY, float64(value))
	}

	return listX, listY

}

func plotGraph(listX, listY []float64) {

	plot, _ := glot.NewPlot(2, false, false)

	points := [][]float64{listX, listY}
	plot.AddPointGroup("Points", "points", points)

	plot.SetTitle("Collatz Conjecture")
	plot.SetXLabel("X-Axis: Range of numbers")
	plot.SetYLabel("Y-Axis: Steps to reach 1")

	maxX := getMax(listX)
	maxY := getMax(listY)

	plot.SetXrange(-1, int((maxX + float64(5))))
	plot.SetYrange(-1, int((maxY + float64(5))))
	plot.SavePlot("github.com/StephenDsouza90/CollatzConjecture/images/Collatz.png")

}

func getMax(listOfNumbers []float64) (max float64) {

	max = listOfNumbers[0]

	for _, value := range listOfNumbers {
		if float64(value) > max {
			max = value
		}
	}

	return max
}
