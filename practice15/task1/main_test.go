package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	input := []int{1, 1, 1, 2, 2, 2}

	expectedOdd := 3
	expectedEven := 3

	actualOdd, actualEven := CalculateEvenAndOddNumbers(input)

	assert.Equal(t, expectedEven, actualEven)
	assert.Equal(t, expectedOdd, actualOdd)
}

func Test_Case2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expectedOdd := 5
	expectedEven := 5

	actualOdd, actualEven := CalculateEvenAndOddNumbers(input)

	assert.Equal(t, expectedEven, actualEven)
	assert.Equal(t, expectedOdd, actualOdd)
}

func Test_Case3(t *testing.T) {
	input := []int{0, 1, 0, 1}

	expectedOdd := 2
	expectedEven := 2

	actualOdd, actualEven := CalculateEvenAndOddNumbers(input)

	assert.Equal(t, expectedEven, actualEven)
	assert.Equal(t, expectedOdd, actualOdd)
}
