package main

func Prepend(array []int) []int {
	newArray := []int{}

	for index := len(array) - 1; index >= 0; index-- {
		newArray = append(newArray, array[index])
	}

	return newArray
}

func ConvertToBinaryRepresentation(number int) []int {
	converted := []int{}
	currentNumber := number

	for currentNumber > 1 {
		remainder := currentNumber % 2
		converted = append([]int{remainder}, converted...)
		currentNumber = currentNumber / 2
	}

	return append([]int{currentNumber}, converted...)
}
