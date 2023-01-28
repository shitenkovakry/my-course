package main

import "fmt"

func MergerArrays(array []int, array2 []int) []int {
	mergedArray := []int{}

	for i := 0; i < len(array); i++ {
		mergedArray = append(mergedArray, array[i])
	}

	for i := 0; i < len(array2); i++ {
		mergedArray = append(mergedArray, array2[i])
	}

	return mergedArray
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	array2 := []int{6, 7, 8, 9, 10}

	result := MergerArrays(array, array2)

	fmt.Println(result)
}
