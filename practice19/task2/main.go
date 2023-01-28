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

func main() {
	array := []int{7, 1}

	result := BubbleSort(array)
	fmt.Println(result)
}
