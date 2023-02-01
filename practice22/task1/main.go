package main

import "fmt"

func FindNumberFromArray(array []int, value int) int {

	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return array[i]
		}
	}

	return value
}

func main() {
	array := []int{5, 9, 1, 2}
	value := 1

	result := FindNumberFromArray(array, value)

	fmt.Println(result)
}
