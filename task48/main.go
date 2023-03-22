package main

func GiveComprasionResult(arrayOfAlise []int, arrayOfBob []int) []int {
	resultOfAliseAndBob := []int{}
	numberOfPointOfAlise := 0
	numberOfPointsOfBob := 0

	for i := 0; i < len(arrayOfAlise); i++ {
		for j := 0; j < len(arrayOfBob); j++ {
			if j == i {
				if arrayOfAlise[i] > arrayOfBob[j] {
					numberOfPointOfAlise += 1
				} else if arrayOfAlise[i] < arrayOfBob[j] {
					numberOfPointsOfBob += 1
				} else if arrayOfAlise[i] == arrayOfBob[j] {
					numberOfPointOfAlise += 0
					numberOfPointsOfBob += 0
				}
			}
		}
	}

	resultOfAliseAndBob = append(resultOfAliseAndBob, numberOfPointOfAlise, numberOfPointsOfBob)

	return resultOfAliseAndBob
}

func main() {

}
