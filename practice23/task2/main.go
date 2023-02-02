package main

import "fmt"

func BubbleSort(array []int) []int {
	bubbledArray := []int{}

	for _, value := range array {
		bubbledArray = append(bubbledArray, value)
	}

	for i := 0; i < len(bubbledArray)-1; i++ {
		start := &bubbledArray[i]
		end := &bubbledArray[i+1]

		if *start > *end {
			*start, *end = *end, *start
			i = -1
		}
	}

	return bubbledArray
}

func Reverse(input []int) []int {
	reversedArray := []int{}

	for i := len(input) - 1; i >= 0; i-- {
		value := input[i]
		reversedArray = append(reversedArray, value)
	}

	return reversedArray
}

func main() {
	x := []int{13, 3, 56, 7, 31, 95, 21, 1}

	y := func(input []int) []int {
		sorted := BubbleSort(input)
		reversed := Reverse(sorted)

		return reversed
	}(x)

	fmt.Println(y)
}
