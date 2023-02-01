package main

import (
	"fmt"
)

func BubbleSort(mergedArray []int) []int {
	bubbledArray := []int{}

	for _, value := range mergedArray {
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

func MergerArrays(array []int, array2 []int) []int {
	mergedArray := []int{}

	for i := 0; i < len(array); i++ {
		mergedArray = append(mergedArray, array[i])
	}

	for i := 0; i < len(array2); i++ {
		mergedArray = append(mergedArray, array2[i])
	}

	result := BubbleSort(mergedArray)

	return result

}

func main() {
	array := []int{1, 5, 8, 9}
	array2 := []int{2, 3, 6, 7, 9}

	result := MergerArrays(array, array2)

	fmt.Println(result)
}
