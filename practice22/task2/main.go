package main

import (
	"errors"
	"fmt"
)

var (
	ErrNoResult = errors.New("нет числа в массиве")
)

func DefineIndex(array []int, value int) (int, error) {

	for i := 0; i < len(array); i++ {
		valueFromArray := array[i]

		if value == valueFromArray {
			return i, nil
		}
	}

	return 0, ErrNoResult
}

func main() {
	array := []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var value int
	fmt.Scan(&value)

	result, err := DefineIndex(array, value)
	if err != nil {
		fmt.Println("нет числа в массиве")

		return
	}

	fmt.Println("индекс равен:", result)
}
