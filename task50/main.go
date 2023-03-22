package main

import "math"

func CalculateSumOfNumbersDiagonally(array [][]int) int {
	summaOfFirst := 0
	summaOfLast := 0

	for i := 0; i < len(array); i++ {
		row := array[i]

		first := row[i]

		summaOfFirst += first

		last := row[len(row)-i-1]

		summaOfLast += last
	}

	resultOfSummaFirstAndLast := int(math.Abs(float64(summaOfFirst - summaOfLast)))

	return resultOfSummaFirstAndLast
}
