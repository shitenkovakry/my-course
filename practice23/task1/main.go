package main

func findMinimalIndexInArray(array []int) int {
	min := array[0]
	minIndex := 0
	for i := 0; i < len(array); i++ {
		value := array[i]

		if value < min {
			min = value
			minIndex = i
		}
	}

	return minIndex
}

func SortByMinimums(input []int) []int {
	result := []int{} // empty array
	array := append([]int{}, input...)

	for {
		if len(array) == 0 {
			break
		}

		minIndex := findMinimalIndexInArray(array)
		value := array[minIndex]
		result = append(result, value)

		left := array[:minIndex]
		right := array[minIndex+1:]

		removed := append(left, right...)

		array = removed
	}

	return result
}

func main() {

}
