package utils

import "math"

// OptimizedGetNumberOfStepsForEachM : This function optimizes the Collatz Conjecture for a given range of numbers
// between startValue and endValue and returns the number of steps to reach 1.
func OptimizedGetNumberOfStepsForEachM(startValue, endValue int) map[int]int {
	//	startValue : starting number of the range
	//	endValue : ending number of the range

	allSteps := make(map[int]int)

	for num := startValue; num <= endValue; num++ {
		countSteps := 0

		m := num
		for m != 1 {
			// Optimization 1:
			//	Stores the steps for each m in the cache, so that the pre-computed value can be reused
			//	next time it comes across the number. This will speed up performance.
			if value, ok := allSteps[m]; ok {
				countSteps = value + countSteps
				break
			}

			//	Optimization 2:
			//	Any number which is a power of 2 will always be even when divided by 2 and will always reach 1.
			//	By taking the log 2 of these numbers, the performance will speed up. The log gives
			//	the number of divisions required to reach 1, which is same as number of steps required to reach 1.
			logOfTwo := math.Log2(float64(m))
			if logOfTwo == float64(int(logOfTwo)) {
				countSteps = int(logOfTwo) + countSteps
				break
			}

			//	Optimization 3:
			//	Since an odd number when multiplied by 3 and added by 1 results to a even number,
			//	it can further be divided by 2 in the same step hence skipping one loop iteration
			//	but still maintaining the steps.
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

	//	allSteps : Dict with m as the key and it's corresponding steps as the value
	return allSteps
}
