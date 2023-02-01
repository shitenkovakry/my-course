package main

import (
	"errors"
	"fmt"
)

var (
	ErrNoResult = errors.New("нет числа в массиве")
)

func FindNumberFromArray(array []int, value int) (int, error) {
	for i := 0; i < len(array); i++ {
		valueFromArray := array[i]

		if valueFromArray == value {
			return valueFromArray, nil
		}
	}

	return 0, ErrNoResult
}

func main() {
	array := []int{5, 9, 1, 2}
	value := 6

	result, err := FindNumberFromArray(array, value)
	if err != nil {
		fmt.Println("нет числа в массиве")

		return
	}

	fmt.Println(result)
}
