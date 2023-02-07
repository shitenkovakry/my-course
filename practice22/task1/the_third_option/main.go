package main

import (
	"errors"
	"fmt"
)

var (
	ErrNoResult = errors.New("нет такого числа")
)

func FindIndexOfValue(input []int, value int) (int, error) {
	for i := 0; i < len(input); i++ {
		if input[i] == value {
			return i, nil
		}
	}

	return 0, ErrNoResult
}

func FindNumbersAfterValue(input []int, value int) ([]int, error) {
	findedNumbers := []int{}

	indexOfValue, err := FindIndexOfValue(input, value)
	if err != nil {
		return nil, ErrNoResult
	}

	for i := 0; i < len(input); i++ {
		if i > indexOfValue {
			findedNumbers = append(findedNumbers, input[i])
		}
	}

	return findedNumbers, nil
}

func CalculateTotalNumbers(input []int) int {
	total := 0

	for i := 1; i < len(input); i++ {
		total += i
	}

	return total
}

func main() {
	input := []int{1, 6, 9, 4, 3, 2}

	fmt.Println("введите число из массива")
	var value int
	fmt.Scan(&value)

	fondNumbers, err := FindNumbersAfterValue(input, value)
	if err != nil {
		fmt.Println("нет результата")
		return
	}

	countedNumbers := CalculateTotalNumbers(fondNumbers)

	fmt.Println(countedNumbers)
}
