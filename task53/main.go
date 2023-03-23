package main

func SearchMin(array []int) int {
	minNumber := array[0]

	for indexOfArray := 0; indexOfArray < len(array); indexOfArray++ {
		if array[indexOfArray] < minNumber {
			minNumber = array[indexOfArray]
		}
	}

	return minNumber
}

func SearchMax(array []int) int {
	maxNumber := array[0]

	for indexOfArray := 0; indexOfArray < len(array); indexOfArray++ {
		if array[indexOfArray] > maxNumber {
			maxNumber = array[indexOfArray]
		}
	}

	return maxNumber
}

func CalculateMaxAndMin(array []int) (int, int) {
	summaOfMax := 0
	summaOfMin := 0

	minNumber := SearchMin(array)
	maxNumber := SearchMax(array)

	for indexOfArray := 0; indexOfArray < len(array); indexOfArray++ {
		summaOfMin += array[indexOfArray]
		summaOfMax += array[indexOfArray]
	}

	return summaOfMin - maxNumber, summaOfMax - minNumber
}
