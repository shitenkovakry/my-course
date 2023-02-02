package main

import "fmt"

func ReturnArraysOfOddAndEven(input []int) ([]int, []int) {
	arrayOfEven := []int{}
	arrayOfOdd := []int{}

	for i := 0; i < len(input); i++ {
		value := input[i]

		if value%2 == 0 {
			arrayOfEven = append(arrayOfEven, value)
		}

		if value%2 != 0 {
			arrayOfOdd = append(arrayOfOdd, value)
		}
	}

	return arrayOfEven, arrayOfOdd
}

func main() {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result, result2 := ReturnArraysOfOddAndEven(input)

	fmt.Println(result)
	fmt.Println(result2)
}
