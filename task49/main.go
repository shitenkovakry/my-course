package main

func CalculateSummaOfArray(array []int) int {
	summaOfArray := 0

	for i := 0; i < len(array); i++ {
		numberFromArray := array[i]
		summaOfArray += numberFromArray
	}

	return summaOfArray
}
