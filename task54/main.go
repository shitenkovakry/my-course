package main

func FindMaxNumberFromArray(array []int) int {
	maxNumber := array[0]

	for indexOfArray := 0; indexOfArray < len(array); indexOfArray++ {
		if array[indexOfArray] > maxNumber {
			maxNumber = array[indexOfArray]
		}
	}

	return maxNumber
}

func CalculateCountOfMaxNumber(array []int) int {
	maxNumber := FindMaxNumberFromArray(array)
	countOfMaxNumber := 0

	for indexOfArray := 0; indexOfArray < len(array); indexOfArray++ {
		if array[indexOfArray] == maxNumber {
			countOfMaxNumber += 1
		}
	}

	return countOfMaxNumber
}
