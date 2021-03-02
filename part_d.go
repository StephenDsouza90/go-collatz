package main

import (
	"fmt"
	"math"
	"time"

	"github.com/Arafatk/glot"
)

func main() {

	var startValue int
	var endValue int

	fmt.Println("Enter start value: ")
	fmt.Scan(&startValue)

	fmt.Println("Enter end value: ")
	fmt.Scan(&endValue)

	start := time.Now()

	allSteps := getNumberOfStepsForEachM(startValue, endValue)

	listX, listY := getAxis(allSteps)

	plotGraph(listX, listY)

	duration := time.Since(start)

	fmt.Printf("\nAll steps for each number within a range of numbers from %v to %v \n\n", startValue, endValue)
	fmt.Println(allSteps, "\n")

	fmt.Printf("--- Took %v seconds ---\n", duration.Seconds())

}

func getNumberOfStepsForEachM(startValue, endValue int) map[int]int {

	allSteps := make(map[int]int)

	for num := startValue; num <= endValue; num++ {

		m := num
		countSteps := 0

		for m != 1 {

			if value, ok := allSteps[m]; ok {
				countSteps = value + countSteps
				break
			}

			logOfTwo := math.Log2(float64(m))
			if logOfTwo == float64(int(logOfTwo)) {
				countSteps = int(logOfTwo) + countSteps
				break
			}

			if m%2 == 0 {
				m = m / 2
				countSteps++
			} else {
				m = (m*3 + 1) / 2
				countSteps += 2
			}

		}

		allSteps[num] = countSteps

	}

	return allSteps
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
	plot.SavePlot("Collatz.png")

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
