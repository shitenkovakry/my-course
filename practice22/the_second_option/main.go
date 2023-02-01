package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())

	var array []int

	for i := 0; i < 10; i++ {
		array = append(array, rand.Intn(10))
	}
	fmt.Println(array)

	var value int
	fmt.Scan(&value)

	result, err := FindNumberFromArray(array, value)
	if err != nil {
		fmt.Println("нет числа в массиве")

		return
	}

	fmt.Println(result)
}
