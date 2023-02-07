package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	input := []int{1, 6, 9, 2, 3, 4}
	value := 9

	expected := []int{2, 3, 4}
	expectedErr := error(nil)
	actual, actualErr := FindNumbersAfterValue(input, value)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedErr, actualErr)
}

func Test_Case2(t *testing.T) {
	input := []int{1, 6, 9, 2, 3, 4}
	value := 9

	expected := 2
	expectedErr := error(nil)
	actual, actualErr := FindIndexOfValue(input, value)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedErr, actualErr)
}

func Test_Case3(t *testing.T) {
	input := []int{2, 3, 4}
	expected := 3

	actual := CalculateTotalNumbers(input)

	assert.Equal(t, expected, actual)
}
