package main

func CalculateDivisionOfNumberOfArrayFromLen(array []int) (float64, float64, float64) {
	totalCountOfPositiveNumbers := 0
	totalCountOfNegativeNumbers := 0
	totalCountOfZero := 0

	for i := 0; i < len(array); i++ {
		numberFromArray := array[i]

		if numberFromArray > 0 {
			totalCountOfPositiveNumbers += 1
		} else if numberFromArray < 0 {
			totalCountOfNegativeNumbers += 1
		} else if numberFromArray == 0 {
			totalCountOfZero += 1
		}

	}

	divisionPositiveNumber := float64(totalCountOfPositiveNumbers) / float64(len(array))
	divisionNegativeNumber := float64(totalCountOfNegativeNumbers) / float64(len(array))
	divisionZero := float64(totalCountOfZero) / float64(len(array))

	return divisionPositiveNumber, divisionZero, divisionNegativeNumber
}

func main() {

}
